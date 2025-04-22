package goju

// Character represents a single Japanese character with its various representations
type Character struct {
	Hiragana string
	Katakana string
	Romaji   string
	Category Category
}

// Category represents the type of character
type Category string

const (
	Seion   Category = "seion"   // 清音
	Dakuon  Category = "dakuon"  // 浊音
	Handaku Category = "handaku" // 半浊音
	Yoon    Category = "yoon"    // 拗音
)

// Characters contains all Japanese characters organized by category
var Characters = map[Category][]Character{
	Seion: {
		{Hiragana: "あ", Katakana: "ア", Romaji: "a"},
		{Hiragana: "い", Katakana: "イ", Romaji: "i"},
		{Hiragana: "う", Katakana: "ウ", Romaji: "u"},
		{Hiragana: "え", Katakana: "エ", Romaji: "e"},
		{Hiragana: "お", Katakana: "オ", Romaji: "o"},
		{Hiragana: "か", Katakana: "カ", Romaji: "ka"},
		{Hiragana: "き", Katakana: "キ", Romaji: "ki"},
		{Hiragana: "く", Katakana: "ク", Romaji: "ku"},
		{Hiragana: "け", Katakana: "ケ", Romaji: "ke"},
		{Hiragana: "こ", Katakana: "コ", Romaji: "ko"},
		{Hiragana: "さ", Katakana: "サ", Romaji: "sa"},
		{Hiragana: "し", Katakana: "シ", Romaji: "shi"},
		{Hiragana: "す", Katakana: "ス", Romaji: "su"},
		{Hiragana: "せ", Katakana: "セ", Romaji: "se"},
		{Hiragana: "そ", Katakana: "ソ", Romaji: "so"},
		{Hiragana: "た", Katakana: "タ", Romaji: "ta"},
		{Hiragana: "ち", Katakana: "チ", Romaji: "chi"},
		{Hiragana: "つ", Katakana: "ツ", Romaji: "tsu"},
		{Hiragana: "て", Katakana: "テ", Romaji: "te"},
		{Hiragana: "と", Katakana: "ト", Romaji: "to"},
		{Hiragana: "な", Katakana: "ナ", Romaji: "na"},
		{Hiragana: "に", Katakana: "ニ", Romaji: "ni"},
		{Hiragana: "ぬ", Katakana: "ヌ", Romaji: "nu"},
		{Hiragana: "ね", Katakana: "ネ", Romaji: "ne"},
		{Hiragana: "の", Katakana: "ノ", Romaji: "no"},
		{Hiragana: "は", Katakana: "ハ", Romaji: "ha"},
		{Hiragana: "ひ", Katakana: "ヒ", Romaji: "hi"},
		{Hiragana: "ふ", Katakana: "フ", Romaji: "fu"},
		{Hiragana: "へ", Katakana: "ヘ", Romaji: "he"},
		{Hiragana: "ほ", Katakana: "ホ", Romaji: "ho"},
		{Hiragana: "ま", Katakana: "マ", Romaji: "ma"},
		{Hiragana: "み", Katakana: "ミ", Romaji: "mi"},
		{Hiragana: "む", Katakana: "ム", Romaji: "mu"},
		{Hiragana: "め", Katakana: "メ", Romaji: "me"},
		{Hiragana: "も", Katakana: "モ", Romaji: "mo"},
		{Hiragana: "や", Katakana: "ヤ", Romaji: "ya"},
		{Hiragana: "ゆ", Katakana: "ユ", Romaji: "yu"},
		{Hiragana: "よ", Katakana: "ヨ", Romaji: "yo"},
		{Hiragana: "ら", Katakana: "ラ", Romaji: "ra"},
		{Hiragana: "り", Katakana: "リ", Romaji: "ri"},
		{Hiragana: "る", Katakana: "ル", Romaji: "ru"},
		{Hiragana: "れ", Katakana: "レ", Romaji: "re"},
		{Hiragana: "ろ", Katakana: "ロ", Romaji: "ro"},
		{Hiragana: "わ", Katakana: "ワ", Romaji: "wa"},
		{Hiragana: "を", Katakana: "ヲ", Romaji: "wo"},
		{Hiragana: "ん", Katakana: "ン", Romaji: "n"},
	},
	Dakuon: {
		{Hiragana: "が", Katakana: "ガ", Romaji: "ga"},
		{Hiragana: "ぎ", Katakana: "ギ", Romaji: "gi"},
		{Hiragana: "ぐ", Katakana: "グ", Romaji: "gu"},
		{Hiragana: "げ", Katakana: "ゲ", Romaji: "ge"},
		{Hiragana: "ご", Katakana: "ゴ", Romaji: "go"},
		{Hiragana: "ざ", Katakana: "ザ", Romaji: "za"},
		{Hiragana: "じ", Katakana: "ジ", Romaji: "ji"},
		{Hiragana: "ず", Katakana: "ズ", Romaji: "zu"},
		{Hiragana: "ぜ", Katakana: "ゼ", Romaji: "ze"},
		{Hiragana: "ぞ", Katakana: "ゾ", Romaji: "zo"},
		{Hiragana: "だ", Katakana: "ダ", Romaji: "da"},
		{Hiragana: "ぢ", Katakana: "ヂ", Romaji: "ji"},
		{Hiragana: "づ", Katakana: "ヅ", Romaji: "zu"},
		{Hiragana: "で", Katakana: "デ", Romaji: "de"},
		{Hiragana: "ど", Katakana: "ド", Romaji: "do"},
		{Hiragana: "ば", Katakana: "バ", Romaji: "ba"},
		{Hiragana: "び", Katakana: "ビ", Romaji: "bi"},
		{Hiragana: "ぶ", Katakana: "ブ", Romaji: "bu"},
		{Hiragana: "べ", Katakana: "ベ", Romaji: "be"},
		{Hiragana: "ぼ", Katakana: "ボ", Romaji: "bo"},
	},
	Handaku: {
		{Hiragana: "ぱ", Katakana: "パ", Romaji: "pa"},
		{Hiragana: "ぴ", Katakana: "ピ", Romaji: "pi"},
		{Hiragana: "ぷ", Katakana: "プ", Romaji: "pu"},
		{Hiragana: "ぺ", Katakana: "ペ", Romaji: "pe"},
		{Hiragana: "ぽ", Katakana: "ポ", Romaji: "po"},
	},
	Yoon: {
		{Hiragana: "きゃ", Katakana: "キャ", Romaji: "kya"},
		{Hiragana: "きゅ", Katakana: "キュ", Romaji: "kyu"},
		{Hiragana: "きょ", Katakana: "キョ", Romaji: "kyo"},
		{Hiragana: "しゃ", Katakana: "シャ", Romaji: "sha"},
		{Hiragana: "しゅ", Katakana: "シュ", Romaji: "shu"},
		{Hiragana: "しょ", Katakana: "ショ", Romaji: "sho"},
		{Hiragana: "ちゃ", Katakana: "チャ", Romaji: "cha"},
		{Hiragana: "ちゅ", Katakana: "チュ", Romaji: "chu"},
		{Hiragana: "ちょ", Katakana: "チョ", Romaji: "cho"},
		{Hiragana: "にゃ", Katakana: "ニャ", Romaji: "nya"},
		{Hiragana: "にゅ", Katakana: "ニュ", Romaji: "nyu"},
		{Hiragana: "にょ", Katakana: "ニョ", Romaji: "nyo"},
		{Hiragana: "ひゃ", Katakana: "ヒャ", Romaji: "hya"},
		{Hiragana: "ひゅ", Katakana: "ヒュ", Romaji: "hyu"},
		{Hiragana: "ひょ", Katakana: "ヒョ", Romaji: "hyo"},
		{Hiragana: "みゃ", Katakana: "ミャ", Romaji: "mya"},
		{Hiragana: "みゅ", Katakana: "ミュ", Romaji: "myu"},
		{Hiragana: "みょ", Katakana: "ミョ", Romaji: "myo"},
		{Hiragana: "りゃ", Katakana: "リャ", Romaji: "rya"},
		{Hiragana: "りゅ", Katakana: "リュ", Romaji: "ryu"},
		{Hiragana: "りょ", Katakana: "リョ", Romaji: "ryo"},
		{Hiragana: "ぎゃ", Katakana: "ギャ", Romaji: "gya"},
		{Hiragana: "ぎゅ", Katakana: "ギュ", Romaji: "gyu"},
		{Hiragana: "ぎょ", Katakana: "ギョ", Romaji: "gyo"},
		{Hiragana: "じゃ", Katakana: "ジャ", Romaji: "ja"},
		{Hiragana: "じゅ", Katakana: "ジュ", Romaji: "ju"},
		{Hiragana: "じょ", Katakana: "ジョ", Romaji: "jo"},
		{Hiragana: "びゃ", Katakana: "ビャ", Romaji: "bya"},
		{Hiragana: "びゅ", Katakana: "ビュ", Romaji: "byu"},
		{Hiragana: "びょ", Katakana: "ビョ", Romaji: "byo"},
		{Hiragana: "ぴゃ", Katakana: "ピャ", Romaji: "pya"},
		{Hiragana: "ぴゅ", Katakana: "ピュ", Romaji: "pyu"},
		{Hiragana: "ぴょ", Katakana: "ピョ", Romaji: "pyo"},
	},
}

// GetCharacterByHiragana returns a character by its hiragana representation
func GetCharacterByHiragana(hiragana string) (Character, bool) {
	for _, chars := range Characters {
		for _, char := range chars {
			if char.Hiragana == hiragana {
				return char, true
			}
		}
	}
	return Character{}, false
}

// GetCharacterByKatakana returns a character by its katakana representation
func GetCharacterByKatakana(katakana string) (Character, bool) {
	for _, chars := range Characters {
		for _, char := range chars {
			if char.Katakana == katakana {
				return char, true
			}
		}
	}
	return Character{}, false
}

// GetCharacterByRomaji returns a character by its romaji representation
func GetCharacterByRomaji(romaji string) (Character, bool) {
	for _, chars := range Characters {
		for _, char := range chars {
			if char.Romaji == romaji {
				return char, true
			}
		}
	}
	return Character{}, false
}
