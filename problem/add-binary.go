// https://leetcode.com/problems/add-binary/
package problem

func AddBinary(a string, b string) string {
	lena, lenb := len(a)-1, len(b)-1
	result := make([]byte, 0)
	carry := 0

	for lena >= 0 || lenb >= 0 || carry > 0 {
		sum := carry

		if lena >= 0 {
			sum += int(a[lena]-'0')
			lena--
		}
		if lenb >= 0 {
			sum += int(b[lenb]-'0')
			lenb--
		}

		result = append(result, byte((sum%2)+'0'))
		carry = sum / 2
	}

	for i, j := 0, len(result)-1; i < j; i , j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
