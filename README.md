# jvmgo

#### Description
golang implement java virtual machine

#### 类文件结构
Class文件是一组以8位字节为基础单位的二进制文件，各个数据项目严格按照顺序紧凑地排列在Class文件中，
中间没有添加任何分隔符，这使得整个Class文件中存储的内容几乎全部都是程序运行的必要数据。
根据Java虚拟机规范，Class文件格式采用一种类似于C语言结构体的伪结构来存储，这种伪结构中只有两种数据类型：无符号数和表。
无符号数属于基本数据类型，以u1、u2、u4、u8来分别代表1、2、4、8个字节的无符号数。
表是由多个无符号数或其他表作为数据项构成的复合数据类型，所有的表都习惯性地以“_info”结尾
整个Class文件本质上就是一张表，它由如下所示的数据项构成。

![01](https://github.com/hadoopeco/jvmgo/blob/master/doc/01.jpg)

从表中可以看出，无论是无符号数还是表，当需要描述同一类型但数量不定的多个数据时，经常会使用一个前置容量计数器加若干个连续数据项的形式，这一系列连续的某类型的数据，称为该数据项的集合，
比如，fields_count个field_info表数据构成了字段表集合。
这里需要说明的是：Class文件中的数据项，都是严格按照上表中的顺序和数量被严格限定的，每个字节代表的含义，长度，先后顺序等都不允许改变


magic与version

每个Class文件的头4个字节称为魔数（magic），它的唯一作用是判断该文件是否为一个能被虚拟机接受的Class文件。它的值固定为0xCAFEBABE。
紧接着magic的4个字节存储的是Class文件的次版本号和主版本号，高版本的JDK能向下兼容低版本的Class文件，但不能运行更高版本的Class文件。

constant_pool

major_version之后是常量池（constant_pool）的入口，它是Class文件中与其他项目关联最多的数据类型，也是占用Class文件空间最大的数据项目之一。

常量池中主要存放两大类常量：字面量和符号引用。字面量比较接近于Java层面的常量概念，
如文本字符串、被声明为final的常量值等。而符号引用总结起来则包括了下面三类常量：

1. 类和接口的全限定名（即带有包名的Class名，如：java.lang.String）
2. 字段的名称和描述符（private、static等描述符）
3. 方法的名称和描述符（private、static等描述符）

虚拟机在加载Class文件时才会进行动态连接，也就是说，Class文件中不会保存各个方法和字段的最终内存布局信息，
因此，这些字段和方法的符号引用不经过转换是无法直接被虚拟机使用的。
当虚拟机运行时，需要从常量池中获得对应的符号引用，再在类加载过程中的解析阶段将其替换为直接引用，并翻译到具体的内存地址中。

这里说明下符号引用和直接引用的区别与关联：

符号引用：符号引用以一组符号来描述所引用的目标，符号可以是任何形式的字面量，只要使用时能无歧义地定位到目标即可。符号引用与虚拟机实现的内存布局无关，引用的目标并不一定已经加载到了内存中。

直接引用：直接引用可以是直接指向目标的指针、相对偏移量或是一个能间接定位到目标的句柄。直接引用是与虚拟机实现的内存布局相关的，同一个符号引用在不同虚拟机实例上翻译出来的直接引用一般不会相同。如果有了直接引用，那说明引用的目标必定已经存在于内存之中了

常量池中的每一项常量都是一个表，共有11种（JDK1.7之前）结构各不相同的表结构数据，没中表开始的第一位是一个u1类型的标志位（1-12，缺少2），代表当前这个常量属于的常量类型。
11种常量类型所代表的具体含义如下表所示：

![02](https://github.com/hadoopeco/jvmgo/blob/master/doc/02.jpg)

这11种常量类型各自均有自己的结构。
在CONSTANT_Class_info型常量的结构中有一项name_index属性，该常属性中存放一个索引值，指向常量池中一个CONSTANT_Utf8_info类型的常量，该常量中即保存了该类的全限定名字符串。
而CONSTANT_Fieldref_info、CONSTANT_Methodref_info、CONSTANT_InterfaceMethodref_info型常量的结构中都有一项index属性，存放该字段或方法所属的类或接口的描述符CONSTANT_Class_info的索引项。
另外，最终保存的诸如Class名、字段名、方法名、修饰符等字符串都是一个CONSTANT_Utf8_info类型的常量，也因此，Java中方法和字段名的最大长度也即是CONSTANT_Utf8_info型常量的最大长度，在CONSTANT_Utf8_info型常量的结构中有一项length属性，它是u2类型的，即占用2个字节，
那么它的最大的length即为65535。因此，Java程序中如果定义了超过64KB英文字符的变量或方法名，将会无法编译。


access_flag

在常量池结束之后，紧接着的2个字节代表访问标志（access_flag），这个标志用于识别一些类或接口层次的访问信息，
包括：这个Class是类还是接口，是否定义为public类型，abstract类型，如果是类的话，是否声明为final，等等。
每种访问信息都由一个十六进制的标志值表示，如果同时具有多种访问信息，则得到的标志值,为这几种访问信息标志值的逻辑或。

this_class、super_class、interfaces

类索引（this_class）和父类索引（super_class）都是一个u2类型的数据，而接口索引集合（interfaces）则是一组u2类型的数据集合，Class文件中由这三项数据来确定这个类的继承关系。
类索引、父类索引和接口索引集合都按照顺序排列在访问标志之后，类索引和父类索引两个u2类型的索引值表示，它们各自指向一个类型为COMNSTANT_Class_info的类描述符常量，通过该常量中的索引值找到定义在COMNSTANT_Utf8_info类型的常量中的全限定名字符串。而接口索引集合就用来描述这个类实现了哪些接口，这些被实现的接口将按implements语句（如果这个类本身是个接口，则应当是extend语句）后的接口顺序从左到右排列在接口的索引集合中。

fields

字段表（field_info）用于描述接口或类中声明的变量。字段包括了类级变量或实例级变量，但不包括在方法内声明的变量。字段的名字、数据类型、修饰符等都是无法固定的，只能引用常量池中的常量来描述。下面是字段表的最种格式：

![03](https://github.com/hadoopeco/jvmgo/blob/master/doc/03.jpg)

其中的access_flags与类中的access_flags类似，是表示数据类型的修饰符，如public、static、volatile等。后面的name_index和descriptor_index都是对常量池的引用，分别代表字段的简单名称及字段和方法的描述符。这里简单解释下“简单名称”、“描述符”和“全限定名”这三种特殊字符串的概念。

前面有所提及，全限定名即指一个事物的完整的名称，如在java.lang包下的String类的全限定名为：java/lang/String，即把包名中的“.”改为“/”，为了使连续的多个全限定名之间不产生混淆，在使用时最后一般会加入一个“，”来表示全限定名结束。
简单名称则是指没有类型或参数修饰的方法或字段名称，如果一个类中有这样一个方法boolean get（int name）和一个变量private final static int m，则他们的简单名称则分别为get（）和m。

而描述符的作用则是用来描述字段的数据类型、方法的参数列表（包括数量、类型以及顺序等）和返回值的。根据描述符规则，详细的描述符标示字的含义如下表所示：

![04](https://github.com/hadoopeco/jvmgo/blob/master/doc/04.jpg)


对于数组类型，每一维度将使用一个前置的“[”字符来描述，如一个整数数组“int [][]”将为记录为“[[I”，
而一个String类型的数组“String[]”将被记录为“[Ljava/lang/String”

用方法描述符描述方法时，按照先参数后返回值的顺序描述，参数要按照严格的顺序放在一组小括号内，
如方法 int getIndex(String name,char[] tgc,int start,int end,char target)的描述符为“（Ljava/lang/String[CIIC）I”。

字段表包含的固定数据项目到descriptor_index为止就结束了，但是在它之后还紧跟着一个属性表集合用于存储一些额外的信息。比如，如果在类中有如下字段的声明：staticfinalint m = 2；那就可能会存在一项名为ConstantValue的属性，它指向常量2。关于attribute_info的详细内容，在后面关于属性表的项目中会有详细介绍。

最后需要注意一点：字段表集合中不会列出从父类或接口中继承而来的字段，但有可能列出原本Java代码中不存在的字段。比如在内部类中为了保持对外部类的访问性，会自动添加指向外部类实例的字段。

methods

方法表（method_info）的结构与属性表的结构相同，不过多赘述。方法里的Java代码，经过编译器编译成字节码指令后，存放在方法属性表集合中一个名为“Code”的属性里，关于属性表的项目，同样会在后面详细介绍。

与字段表集合相对应，如果父类方法在子类中没有被覆写，方法表集合中就不会出现来自父类的方法信息。
但同样，有可能会出现由编译器自动添加的方法，最典型的便是类构造器“<clinit>”方法和实例构造器“<init>”方法。

在Java语言中，要重载一个方法，除了要与原方法具有相同的简单名称外，还要求必须拥有一个与原方法不同的特征签名，特征签名就是一个方法中各个参数在常量池中的字段符号引用的集合，也就是因为返回值不会包含在特征签名之中，因此Java语言里无法仅仅依靠返回值的不同来对一个已有方法进行重载。

attributes

属性表（attribute_info）在前面已经出现过多次，在Class文件、字段表、方法表中都可以携带自己的属性表集合，以用于描述某些场景专有的信息。

属性表集合的限制没有那么严格，不再要求各个属性表具有严格的顺序，并且只要不与已有的属性名重复，任何人实现的编译器都可以向属性表中写入自己定义的属性信息，但Java虚拟机运行时会忽略掉它不认识的属性。
Java虚拟机规范中预定义了9项虚拟机应当能识别的属性（JDK1.5后又增加了一些新的特性，因此不止下面9项，但下面9项是最基本也是必要，出现频率最高的），如下表所示：

![05](https://github.com/hadoopeco/jvmgo/blob/master/doc/05.jpg)

对于每个属性，它的名称都需要从常量池中引用一个CONSTANT_Utf8_info类型的常量来表示，每个属性值的结构是完全可以自定义的，只需说明属性值所占用的位数长度即可。一个符合规则的属性表至少应具有“attribute_name_info”、“attribute_length”和至少一项信息属性。

1）Code属性

前面已经说过，Java程序方法体中的代码讲过Javac编译后，生成的字节码指令便会存储在Code属性中，但并非所有的方法表都必须存在这个属性，比如接口或抽象类中的方法就不存在Code属性。如果方法表有Code属性存在，那么它的结构将如下表所示：

![06](https://github.com/hadoopeco/jvmgo/blob/master/doc/06.jpg)


attribute_name_index是一项指向CONSTANT_Utf8_info型常量的索引，常量值固定为“Code”，它代表了该属性的名称。attribute_length指示了属性值的长度，由于属性名称索引与属性长度一共是6个字节，所以属性值的长度固定为整个属性表的长度减去6个字节。

max_stack代表了操作数栈深度的最大值，max_locals代表了局部变量表所需的存储空间，它的单位是Slot，并不是在方法中用到了多少个局部变量，就把这些局部变量所占Slot之和作为max_locals的值，原因是局部变量表中的Slot可以重用。

code_length和code用来存储Java源程序编译后生成的字节码指令。code用于存储字节码指令的一系列字节流，它是u1类型的单字节，因此取值范围为0x00到0xFF，那么一共可以表达256条指令，目前，Java虚拟机规范已经定义了其中200条编码值对应的指令含义。
code_length虽然是一个u4类型的长度值，理论上可以达到2^32-1，但是虚拟机规范中限制了一个方法不允许超过65535条字节码指令，如果超过了这个限制，Javac编译器将会拒绝编译。

字节码指令之后是这个方法的显式异常处理表集合（exception_table），它对于Code属性来说并不是必须存在的。它的格式如下表所示：

![07](https://github.com/hadoopeco/jvmgo/blob/master/doc/07.jpg)

它包含四个字段，这些字段的含义为：如果字节码从第start_pc行到第end_pc行之间（不含end_pc行）出现了类型为catch_type或其子类的异常（catch_type为指向一个CONSTANT_Class_info型常量的索引），则转到第handler_pc行继续处理，当catch_pc的值为0时，代表人和的异常情况都要转到handler_pc处进行处理。异常表实际上是Java代码的一部分，编译器使用异常表而不是简单的跳转命令来实现Java异常即finally处理机制，也因此，finally中的内容会在try或catch中的return语句之前执行，并且在try或catch跳转到finally之前，会将其内部需要返回的变量的值复制一份副本到最后一个本地表量表的Slot中，也因此便有了http://blog.csdn.net/ns_code/article/details/17485221这篇文章中出现的情况。

1）Code属性

Code属性是Class文件中最重要的一个属性，如果把一个Java程序中的信息分为代码和元数据两部分，那么在整个Class文件里，Code属性用于描述代码，所有的其他数据项目都用于描述元数据。

2）Exception属性

这里的Exception属性的作用是列举出方法中可能抛出的受查异常，也就是方法描述时在throws关键字后面列举的异常。它的结构很简单，只有attribute_name_index、attribute_length、number_of_exceptions、exception_index_table四项，从字面上便很容易理解，这里不再详述。

3）LineNumberTable属性

它用于描述Java源码行号与字节码行号之间的对应关系。

4）LocalVariableTable属性

它用于描述栈帧中局部变量表中的变量与Java源码中定义的变量之间的对应关系。

5）SourceFile属性

它用于记录生成这个Class文件的源码文件名称。

6）ConstantValue属性

ConstantValue属性的作用是通知虚拟机自动为静态变量赋值，只有被static修饰的变量才可以使用这项属性。在Java中，对非static类型的变量（也就是实例变量）的赋值是在实例构造器<init>方法中进行的；而对于类变量（static变量），则有两种方式可以选择：在类构造其中赋值，或使用ConstantValue属性赋值。

目前Sun Javac编译器的选择是：如果同时使用final和static修饰一个变量（即全局常量），并且这个变量的数据类型是基本类型或String的话，就生成ConstantValue属性来进行初始化（编译时Javac将会为该常量生成ConstantValue属性，在类加载的准备阶段虚拟机便会根据ConstantValue为常量设置相应的值），如果该变量没有被final修饰，或者并非基本类型及字符串，则选择在<clinit>方法中进行初始化。

虽然有final关键字才更符合”ConstantValue“的含义，但在虚拟机规范中并没有强制要求字段必须用final修饰，只要求了字段必须用static修饰，对final关键字的要求是Javac编译器自己加入的限制。因此，在实际的程序中，只有同时被final和static修饰的字段才有ConstantValue属性。而且ConstantValue的属性值只限于基本类型和String，很明显这是因为它从常量池中也只能够引用到基本类型和String类型的字面量。

下面简要说明下final、static、static final修饰的字段赋值的区别：

static修饰的字段在类加载过程中的准备阶段被初始化为0或null等默认值，而后在初始化阶段（触发类构造器<clinit>）才会被赋予代码中设定的值，如果没有设定值，那么它的值就为默认值。

final修饰的字段在运行时被初始化（可以直接赋值，也可以在实例构造器中赋值），一旦赋值便不可更改；

static final修饰的字段在Javac时生成ConstantValue属性，在类加载的准备阶段根据ConstantValue的值为该字段赋值，它没有默认值，必须显式地赋值，否则Javac时会报错。可以理解为在编译期即把结果放入了常量池中。

7）InnerClasses属性

该属性用于记录内部类与宿主类之间的关联。如果一个类中定义了内部类，那么编译器将会为它及它所包含的内部类生成InnerClasses属性。

8）Deprecated属性

该属性用于表示某个类、字段和方法，已经被程序作者定为不再推荐使用，它可以通过在代码中使用@Deprecated注释进行设置。

9）Synthetic属性

该属性代表此字段或方法并不是Java源代码直接生成的，而是由编译器自行添加的，如this字段和实例构造器、类构造器等。

#### 

1.  go install .

2. run 
jvmgo.exe  -Xjre "$JAVA_HOME"  Test




