# 问题1：两个任意位一进制数加法
输入:(若干1)+(若干1)=

停机时纸带:...###(若干1)B###...

其中输入的“若干1”表示要做加法的一进制数，可能是零个1。

停机时纸带上的“若干1” 作为答案，它的起始位置应为输入的 “=” 位置右侧一格，并以 B 作为结尾。# 为 Γ 中任意字母。

## 示例:

输入: 11+1=

停机时纸带: ...11+1=111B...

输入: 11+1=

停机时纸带: ...FBc01111BcWW... (注: 输出中蓝色的“1”和输入中“=”的位置相同)

# 问题2: 两个4位二进制数加法

输入: (4位0或者1)+(4位0或者1)=

停机时纸带: ...###(5位0或者1)B###...

输入的“4位0或者1”表示要做加法的二进制数

停机时纸带上的“5位0或者1” （结果只有4位时注意加前导0）作为答案，起始位置应为输入中“=” 位置右侧一格，并以 B 作为结尾。# 为 Γ 中任意字母。

## 样例:

输入: 1101+0011=

停机时纸带: ...1101+0011=10000B...

输入: 0001+0011=

停机时纸带: ...001101101 **0** 00100B000... (注: 输出中加粗的“0”和输入中“=”的位置相同)

# 问题3: 两个任意位二进制数加法

输入: ((若干0或者1)+(若干0或者1)=

停机时纸带: ...###(若干0或者1)B###...

输入的“若干0或者1”中不含前导0；每一个操作数的长度大于等于1。

T停机时纸带上的“若干0或者1”作为答案，不应含有前导0。答案的起始位置应为输入中 “=” 位置右侧一格，并以 B 作为结尾。# 为 Γ 中任意字母。

## 样例:

输入:1+1001=

停机时纸带: ...11+1001=1100B...

输入:11+1001=

停机时纸带: ...wRQTqLaA1100BssR...（注: 输出中的“A”和输入中“=”的位置相同）
