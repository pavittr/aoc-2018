package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {

	checksumCalc := func(sequence string) int {
		countOf3s := 0
		countOf2s := 0
		for _, line := range strings.Split(sequence, "\n") {
			checkSum := map[rune]int{}
			for _, char := range line {
				checkSum[char] += 1
			}
			matches2 := false
			matches3 := false
			for _, counters := range checkSum {
				if counters == 3 {
					matches3 = true
				}
				if counters == 2 {
					matches2 = true
				}
			}
			if matches3 {
				countOf3s++
			}
			if matches2 {
				countOf2s++
			}
		}
		return countOf3s * countOf2s
	}
	testInput := `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	assert.Equal(t, 12, checksumCalc(testInput))
	assert.Equal(t, 5928, checksumCalc(day2Input))

}

func TestDay2Part2(t *testing.T) {
	diffCalc := func(sequence string) string {
		list := make([]string, 0)
		for _, line := range strings.Split(sequence, "\n") {
			list = append(list, line)
		}
		for _, outerWord := range list {
			for _, innerWord := range list {
				resultWord := make([]byte, 0)
				if innerWord == outerWord {
					continue
				}
				for i := 0; i < len(outerWord); i++ {
					if innerWord[i] == outerWord[i] {
						resultWord = append(resultWord, innerWord[i])
					}
				}
				if len(resultWord) == len(outerWord)-1 {
					return string(resultWord)
				}
			}
		}
		return ""
	}
	testInput := `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

	assert.Equal(t, "fgij", diffCalc(testInput))
	assert.Equal(t, "bqlporuexkwzyabnmgjqctvfs", diffCalc(day2Input))
}

var day2Input = `bqlpzruexkszyahnamgjdctvfs
bqlporuexkwyyahnbmgjdctvfb
bqlhoruexkwzyahefmgjdctvfs
bqlporuegkwzyahdimgodctvfs
bqlporuxxkwhybhnimgjdctvfs
dqlporuexkwzyahnimgjfcivfs
bqlporlexkwzyarniigjdctvfs
bqlporuezkqzeahnimgjdctvfs
bqlporuefkwzyamnsmgjdctvfs
bqlporuexkwzyahqimgjdazvfs
bqspocuexkwzyahnimgjdctffs
bqlporuexkwzywhnimgjiccvfs
bqlcoruexkwzyahnimgjdgtvfj
bqljoruexkwvyahnimgcdctvfs
bqeporuexkwzyahnionjdctvfs
bqlporuexkwzyabnxmgjqctvfs
bqjporuexkwzyahnimtddctvfs
bqlprruexkezyahbimgjdctvfs
bqlporuexkazyahnimgjjctefs
bqlioruexkweyjhnimgjdctvfs
bqlporuexsszymhnimgjdctvfs
bqlporffxkwzyahnimgjditvfs
bqlporuexkwzyahnimcedctvbs
bqwpgruexkezyahnimgjdctvfs
bqldoruexkwzyahnimgjdcpvfu
bqlpoxueckwzyahnemgjdctvfs
bqlhoruqxkwzyahnimgjdctsfs
wqlporuexkwzychfimgjdctvfs
bqlporuenkwzyahnimvqdctvfs
aqlppruexkwzyahnimgjdcrvfs
bqlporuexzwzfahnimgjdcivfs
bqjporuexkwzyahnimgjdcqmfs
bqppkkuexkwzyahnimgjdctvfs
bqlporuervwzyahnimgldctvfs
bqoporxeykwzyahnimgjdctvfs
bqyporuehkwzyahnilgjdctvfs
bqlporuexkwzybvnimgjdctvfv
bqlporzehkwzyaznimgjdctvfs
bqiporuexkjzyahnimgjdatvfs
yqlpkruexkwzyahnimggdctvfs
brlporuexkwzyehnimgjdctvbs
bqlporuyxkwzyahnapgjdctvfs
bqlporuexkwzyaewimgjdttvfs
bqlporueskwzyahnimpjdcdvfs
bqlioruexkwzyahnimghdcbvfs
bqljoruexkwzyahtimgjdctmfs
bqlporgexkwzyaxnimgjvctvfs
bqlporuexkwxyahnimgodftvfs
mqlporupxkwzyahnimgjdcqvfs
bgkporuexkwzychnimgjdctvfs
jqlporuexgwzyawnimgjdctvfs
bqlperuebkwzyahniwgjdctvfs
bqlporuexkwzyahnbmgjdntvts
bqlporuexkwzyajnimgldctvfz
bqlpobuexkwzychnimgjdctvfu
bquporuexkwzyahiimgjdctvfp
bnlporbexkwzyahnimgjdctvfb
bqlporuexklzyahnimgtdctbfs
bqlworuexkwzyahnimkjdcpvfs
bqlporumxkwzyahnikejdctvfs
bqlporuexkwryahnimfjdctnfs
bqlioruenkwzywhnimgjdctvfs
bqlporulxbwzyahnimgjdctvfe
bqeporufxkwzqahnimgjdctvfs
bqlpbwuexkwzyahnimgjdctvfo
bqlpoduexkwzyahnimgvdctvrs
bqlporuexkwzlahnimgjdctvdw
bqlporuexkwzyahnimujdctdfp
iqlporuexvwvyahnimgjdctvfs
bclporuexkwzyahnimgjdzovfs
bqlporuerktzyahnamgjdctvfs
bqlporuexkwziajnimgjdctvfe
bqlpnrkexkwoyahnimgjdctvfs
bqlporuexkwznacnimgjdctvks
bqlporuexkrzyrhnimgjdcuvfs
bqltoruexkwzyahnnmgjdcivfs
bqlporuexkwzuahnimdjdctmfs
bqlporubxkbzyahnimgjdctmfs
bqlplruexkwdyhhnimgjdctvfs
bylporuexknzyahnimgjdctmfs
bqlporuexkwzyahqitgjdctvas
bquperuexkwzyahnimgjdcdvfs
bqlporuexktzyahnisgjdctvls
bqlporuexkwfyacnimtjdctvfs
bqleoraexkwzyahnimgjoctvfs
bqlporuexkwlyahnimrjdctvas
bqlphruexkwzyadnimujdctvfs
oqadoruexkwzyahnimgjdctvfs
belpzruexkwzyahnimgjjctvfs
bqlporuexkwzkahtbmgjdctvfs
bqlporulctwzyahnimgjdctvfs
bklptrwexkwzyahnimgjdctvfs
bqlpoqhexkwzjahnimgjdctvfs
tqlporjexkwzyahnimgjdctvfx
bqipwruexkwzyahnimgjdctvfd
bqlparueeawzyahnimgjdctvfs
bqlsoruexkrzyahnqmgjdctvfs
bqlsoruexhwzwahnimgjdctvfs
bquporuexkwsyihnimgjdctvfs
bwlporuexkwzyahwimgjdcuvfs
bqgpobuexkczyahnimgjdctvfs
bqlporuexkwzyahntmgjdspvfs
bqlporuetkwzjavnimgjdctvfs
bqlpoluwxkwzyahnimgpdctvfs
bqlporueykwzuahnimgjuctvfs
bqqporuexkwzyranimgjdctvfs
bjlporuexkwzyahnjmgjdctvts
bjlpofuenkwzyahnimgjdctvfs
bqhporuexbwzyahnfmgjdctvfs
bqlpofulxkwzyfhnimgjdctvfs
bqlporuexrwzyahuimgjdcwvfs
bblporaexkwzyahnicgjdctvfs
uqlpoxuexkwzuahnimgjdctvfs
bqlporuexkwzylhnimgjjhtvfs
bqlloruexkhiyahnimgjdctvfs
bqlpopuexkwzyahnymojdctvfs
bqlporuexkwgyyhnimgjdstvfs
bqlroruexxwzyahnumgjdctvfs
balpopuexkwzeahnimgjdctvfs
bqlporuexkwzyahniogjdqtvfc
bqlpoiuexkwgynhnimgjdctvfs
bqlpoyueukwzyahnimgjdcwvfs
bqtporuexkwtyahnimgjdctvfl
bqlporzexkwzyahnsmgjdctxfs
bqlporukxkwhyahnimgddctvfs
bklporuevowzyahnimgjdctvfs
bqgporuexmwzyahnimgjdctsfs
bqlporuetkxzyahnimgjoctvfs
bqlpsrzexkwzyahnimcjdctvfs
bqlporuexkwzzahaimgkdctvfs
bqlporuexkwzyahnimgjdmtvwd
bqlporuexkwzywhlimgjhctvfs
bqlporuexkwuhahnimgjdctvns
bqqporuexkizyahnimgjdctcfs
bqlvoruepkbzyahnimgjdctvfs
bqlporuexkwqyahnimgjfctvbs
bqlporuexkwzyahlimghuctvfs
bqlporuexkwzyahnizgjictvfb
bqlporuvxkjzyahaimgjdctvfs
bqljoruexkwzyahnizgjdctqfs
bqlporuexkwyyahnimgjdztvfv
jqlporqexkwzyahvimgjdctvfs
bqlporueakwzyabnimgjdctxfs
bqlporuezewzyahnimgfdctvfs
bqlporuzxklzyahnwmgjdctvfs
zqlpproexkwzyahnimgjdctvfs
bqlporuefkwzyahnlmgjdgtvfs
bqlporoenkwzyahnkmgjdctvfs
bqwporuexkwzyahnimcjdctvfk
bqlporuexkwbyahaimgjdctvus
bqlporujxkwzyahnimgjfntvfs
bqlooruexkvzyahnimjjdctvfs
bqlpomuexowzyahnimgjcctvfs
bqlpoquexkwzvahnimgjdctvfl
bclpopuexowzyahnimgjdctvfs
bqlporuexcwzyahnimgjdctlfb
bqlporgexkwzyalnimgjdctvfq
bqppiruexkwzyyhnimgjdctvfs
bqlporaexkwnmahnimgjdctvfs
bqhporuexkwzyahlomgjdctvfs
bqlloruexkwzuahnimgjdctvfb
bqhporuexkwzyahnemgjdcovfs
bqlpiruexkwdyahwimgjdctvfs
bqlporuexkwryahnimgjtdtvfs
bqeporuexkwzyahgixgjdctvfs
bqlporuexkwyqahnimgfdctvfs
bqlporuexkwzyaoniggydctvfs
bqlpouuexkazyahnimgxdctvfs
brlporuexkwzyahvimgjdctvps
bqlforuexkwzyrhnitgjdctvfs
bqlpiruexkwzyfhnjmgjdctvfs
bqlpoebexkwzyxhnimgjdctvfs
bblporoezkwzyahnimgjdctvfs
bglporuexkwzyajqimgjdctvfs
bqlpxcuixkwzyahnimgjdctvfs
bqlporuaekwzyahniegjdctvfs
hqvporuexkwzyahnimgjectvfs
qllporuexkwzyahnimgjdctvhs
bqlporaexkpzyahnlmgjdctvfs
bqlporuexkwzyabnzmgjqctvfs
bqlporuexbvzyahnimgjdctrfs
bqdzoruetkwzyahnimgjdctvfs
bslporuexkwzyahnimgjdctdos
bqlporfexkwzylwnimgjdctvfs
bqrporueykwzyavnimgjdctvfs
bqlporuexkweythnomgjdctvfs
bqlpozuepkwzyahnimgjdctdfs
bqoporuexqhzyahnimgjdctvfs
bqlporucxkwzyahnimtjdctifs
bqlpobuexawzyahnimgjdcthfs
dklporuexkwzyahnimhjdctvfs
bclpkquexkwzyahnimgjdctvfs
bqlporuexkwzyafnfmljdctvfs
bqlporuexkwzympnimgjsctvfs
bqlporuexkwzyaonimdjactvfs
bqlporuvxkwzywhnimgjdctvks
bqlporuexkwzgahnimbjdctvfn
bqlqoruexvwzqahnimgjdctvfs
bqlporuexkmkyahniggjdctvfs
bqlneruexkwzyjhnimgjdctvfs
bqlporueqkwzyahnbmgjdctvfl
bqlpgruexkwzyahnimjvdctvfs
nqlpsruevkwzyahnimgjdctvfs
bqhpouuexkwzyahnmmgjdctvfs
bqnporuexkwzyahnimojdctqfs
bqlpordexkwztahnimgjdctvds
eqlporufxkwzyahniigjdctvfs
balporuebkwzyahnimgjdctvfv
bqlprruexrwzwahnimgjdctvfs
bqlporuepkwxyahnimggdctvfs
bqlporfexkwzyahnqjgjdctvfs
bqaporuexvwzyahnimgjdctvfy
bqlporuexkwzyawnibgjdctmfs
eqlhoruexkwzyahiimgjdctvfs
bqlporuexkbzyahnimgsdhtvfs
bqlporhexkfzyahnimgjdcgvfs
bqeporuexkxzyahnimgjdcavfs
bqlporoerkwzyahnimgjdctnfs
bqlporuemkwzyshnimgjdatvfs
bilporuexkwyyahnimgvdctvfs
xqlporuexkwzyahnnmgjdctvfl
bqlborumxkwzsahnimgjdctvfs
bqlporsexcwzvahnimgjdctvfs
bqlporxexkwzyapnimgbdctvfs
bqlpmruexkwzyahnimgbdcpvfs
qqlroruexkwzyahnihgjdctvfs
bqnporuexkbzyaknimgjdctvfs
bqxpoeuexcwzyahnimgjdctvfs
bqlporuexkqzyahnidgjdcivfs
zqkmoruexkwzyahnimgjdctvfs
bqlporuexkwzyahzimgjqjtvfs
bqfporuenkwzyahiimgjdctvfs
bqlporuexkwzyahuimszdctvfs
bklporurmkwzyahnimgjdctvfs
bqlpotuexktzyahnimgjdcfvfs
bqlporuexkwxyahnimgxdltvfs
bqmporuexkwzoahnimgjdctvys
bqlporuexvwzyahnimgjdnnvfs
bqzporuexkwzyahnkmgjdctsfs
bqmporuexkwzyahnihgjdctvfq
bqlporuexkwzyavnimgjdcrvos
bklpopuexkwzyahnimgjdctvfb
bvlcoruexkwzyapnimgjdctvfs
bqlboruexzwcyahnimgjdctvfs
bqlporuexdwzyihnimgydctvfs
bqlpovuexkwzyaynimgjdctvfq
bqlarruebkwzyahnimgjdctvfs
fqlporuexkwzfahnimgjdctsfs
bqlporuexowzyahnjmgjdctdfs
bqlporuexkwzyahnrmkjdctyfs`
