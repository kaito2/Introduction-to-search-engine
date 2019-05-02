package modules


var ignoredChars = []rune{
	' ',
	'\f',
	'\n',
	'\r',
	'\t',
	'\v',
	'!',
	'"',
	'#',
	'$',
	'%',
	'&',
	'\'',
	'(',
	')',
	'*',
	'+',
	',',
	'-',
	'.',
	'/',
	':',
	';',
	'<',
	'=',
	'>',
	'?',
	'@',
	'[',
	'\\',
	']',
	'^',
	'_',
	'`',
	'{',
	'|',
	'}',
	'~',
	0x3000,  // Double-byte space
	0x3001,  // '、'
	0x3002,  // '。'
	0xFF08,  // '（'
	0xFF09,  // '）'
}


func IsIgnoredChar(ustr rune) bool {
	for _, ignoredChar := range ignoredChars {
		if ustr == ignoredChar {
			return true
		}
	}
	return false
}
