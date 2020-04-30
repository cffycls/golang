package main

import (
	"fmt"
	"strconv"
	"strings"
)

/** java
public class Solution {
    public String mirroring(String s) {
        String x = s.substring(0, (s.length()) / 2);
        return x + (s.length() % 2 == 1 ? s.charAt(s.length() / 2) : "") + new StringBuilder(x).reverse().toString();
    }
    public String nearestPalindromic(String n) {
        if (n.equals("1"))
            return "0";

        String a = mirroring(n);
        long diff1 = Long.MAX_VALUE;
        diff1 = Math.abs(Long.parseLong(n) - Long.parseLong(a));
        if (diff1 == 0)
            diff1 = Long.MAX_VALUE;

        StringBuilder s = new StringBuilder(n);
        int i = (s.length() - 1) / 2;
        while (i >= 0 && s.charAt(i) == '0') {
            s.replace(i, i + 1, "9");
            i--;
        }
        if (i == 0 && s.charAt(i) == '1') {
            s.delete(0, 1);
            int mid = (s.length() - 1) / 2;
            s.replace(mid, mid + 1, "9");
        } else
            s.replace(i, i + 1, "" + (char)(s.charAt(i) - 1));
        String b = mirroring(s.toString());
        long diff2 = Math.abs(Long.parseLong(n) - Long.parseLong(b));


        s = new StringBuilder(n);
        i = (s.length() - 1) / 2;
        while (i >= 0 && s.charAt(i) == '9') {
            s.replace(i, i + 1, "0");
            i--;
        }
        if (i < 0) {
            s.insert(0, "1");
        } else
            s.replace(i, i + 1, "" + (char)(s.charAt(i) + 1));
        String c = mirroring(s.toString());
        long diff3 = Math.abs(Long.parseLong(n) - Long.parseLong(c));

        if (diff2 <= diff1 && diff2 <= diff3)
            return b;
        if (diff1 <= diff3 && diff1 <= diff2)
            return a;
        else
            return c;
    }
}

作者：LeetCode
链接：https://leetcode-cn.com/problems/find-the-closest-palindrome/solution/xun-zhao-zui-jin-de-hui-wen-shu-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func main()  {
	n := "123456789"
	s := []byte(n)
	nv,_ := strconv.Atoi(n)
	fmt.Println(s, len(s)/2, strings.Split(n,""), nv)
	/*for _,v := range s {
		fmt.Print(string([] byte{v})
	}*/

	fmt.Print(nearestPalindromic("1001"))


}

func nearestPalindromic(n string) string {
	//nv,_ := strconv.Atoi(n)
	s := strings.Split(n, "")
	val,_ := strconv.Atoi(n)
	pos1, pos2, num := 1, 1, 0
	//前一半数
	fmt.Println("*****")
	midv := 0
	for i:=0; (i+1)*2<=len(s)+1; i++ {
		fmt.Println(i, (i+1)*2, len(s)+1)
		sv,_ := strconv.Atoi(s[i])
		pos1 = 1
		for j:=1; j<len(s)-i; j++ {
			pos1 *= 10
		}
		pos2 = 1
		for j:=0; j<i; j++ {
			pos2 *= 10
		}
		if pos1 >= pos2 {
			num += sv * pos1
			if pos1 != pos2{
				num += sv * pos2
				//fmt.Println("num = ", num)
			}
			fmt.Println("num** = ", num, ";pos1= ", pos1, ";pos2=", pos2, ";midv=", sv, "; 中卫 i=", i)
		}
	}
	fmt.Println("等于自己时，要往前找个回文数")
	//1.等于自己时，要往前找个回文数
	more1, more2 := 0, 0
	if pos1 == pos2 { //偶数个
		pre := val
		for pre>=10 {
			pre /= 10
		}
		fmt.Println("....", midv, pre)
		if (midv==0) && (pre==1) { //借位，位数减少了
			more1 = num - 2
			more2 = num + (pos1 + pos2)
		}else{
			more1 = num - pos1
			more2 = num + pos1
		}
	} else { //奇数个
		pre := val
		for pre>=10 {
			pre /= 10
		}
		fmt.Println("-----", midv, pre)
		if midv==0 { //借位，位数减少了
			if pre>1 || 2*(pos1*10)<num {
				more1 = num - (pos1*10) + pos1 - (pos1/10)
			}else{

			}
			more2 = num + (pos1 + pos2)
			fmt.Println("-----", more1, more2)

		}else{
			more1 = num - (pos1 + pos2)
			more2 = num + (pos1 + pos2)
		}
	}
	//等于自身等的情况
	// more1-num-more2
	fmt.Println("val** = ", val, ";more1= ", more1, ";num=", num, "; more2=", more2)
	if (val-more1) <= (more2-val) {
		num = more1
	} else {
		num = more2
	}
	return strconv.Itoa(num)
}