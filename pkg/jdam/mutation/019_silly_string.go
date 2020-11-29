package mutation

import (
	"math/rand"
	"reflect"
)

var sillyStringPayloads = []string{
	"NaN",
	"Infinity",
	"-Infinity",
	"undefined",
	"undef",
	"null",
	"NULL",
	"(null)",
	"nil",
	"NIL",
	"true",
	"false",
	"True",
	"False",
	"TRUE",
	"FALSE",
	"None",
	"hasOwnProperty",
	"\\",
	"\\\\",
	",./;'[]\\-=",
	"<>?:\"{}|_+",
	"!@#$%^&*()`~",
	"Ω≈ç√∫˜µ≤≥÷",
	"åß∂ƒ©˙∆˚¬…æ",
	"œ∑´®†¥¨ˆøπ“‘",
	"¡™£¢∞§¶•ªº–≠",
	"¸˛Ç◊ı˜Â¯˘¿",
	"ÅÍÎÏ˝ÓÔÒÚÆ☃",
	"Œ„´‰ˇÁ¨ˆØ∏”’",
	"`⁄€‹›ﬁﬂ‡°·‚—±",
	"⅛⅜⅝⅞",
	"ЁЂЃЄЅІЇЈЉЊЋЌЍЎЏАБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдежзийклмнопрстуфхцчшщъыьэюя",
	"٠١٢٣٤٥٦٧٨٩",
	"⁰⁴⁵",
	"₀₁₂",
	"⁰⁴⁵₀₁₂",
	"ด้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็ ด้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็ ด้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็็้้้้้้้้็็็็็้้้้้็็็็",
	"田中さんにあげて下さい",
	"パーティーへ行かないか",
	"和製漢語",
	"部落格",
	"사회과학원 어학연구소",
	"찦차를 타고 온 펲시맨과 쑛다리 똠방각하",
	"社會科學院語學研究所",
	"울란바토르",
	"𠜎𠜱𠝹𠱓𠱸𠲖𠳏",
	"Ｔｈｅ ｑｕｉｃｋ ｂｒｏｗｎ ｆｏｘ ｊｕｍｐｓ ｏｖｅｒ ｔｈｅ ｌａｚｙ ｄｏｇ",
	"𝐓𝐡𝐞 𝐪𝐮𝐢𝐜𝐤 𝐛𝐫𝐨𝐰𝐧 𝐟𝐨𝐱 𝐣𝐮𝐦𝐩𝐬 𝐨𝐯𝐞𝐫 𝐭𝐡𝐞 𝐥𝐚𝐳𝐲 𝐝𝐨𝐠",
	"𝕿𝖍𝖊 𝖖𝖚𝖎𝖈𝖐 𝖇𝖗𝖔𝖜𝖓 𝖋𝖔𝖝 𝖏𝖚𝖒𝖕𝖘 𝖔𝖛𝖊𝖗 𝖙𝖍𝖊 𝖑𝖆𝖟𝖞 𝖉𝖔𝖌",
	"𝑻𝒉𝒆 𝒒𝒖𝒊𝒄𝒌 𝒃𝒓𝒐𝒘𝒏 𝒇𝒐𝒙 𝒋𝒖𝒎𝒑𝒔 𝒐𝒗𝒆𝒓 𝒕𝒉𝒆 𝒍𝒂𝒛𝒚 𝒅𝒐𝒈",
	"𝓣𝓱𝓮 𝓺𝓾𝓲𝓬𝓴 𝓫𝓻𝓸𝔀𝓷 𝓯𝓸𝔁 𝓳𝓾𝓶𝓹𝓼 𝓸𝓿𝓮𝓻 𝓽𝓱𝓮 𝓵𝓪𝔃𝔂 𝓭𝓸𝓰",
	"𝕋𝕙𝕖 𝕢𝕦𝕚𝕔𝕜 𝕓𝕣𝕠𝕨𝕟 𝕗𝕠𝕩 𝕛𝕦𝕞𝕡𝕤 𝕠𝕧𝕖𝕣 𝕥𝕙𝕖 𝕝𝕒𝕫𝕪 𝕕𝕠𝕘",
	"𝚃𝚑𝚎 𝚚𝚞𝚒𝚌𝚔 𝚋𝚛𝚘𝚠𝚗 𝚏𝚘𝚡 𝚓𝚞𝚖𝚙𝚜 𝚘𝚟𝚎𝚛 𝚝𝚑𝚎 𝚕𝚊𝚣𝚢 𝚍𝚘𝚐",
	"⒯⒣⒠ ⒬⒰⒤⒞⒦ ⒝⒭⒪⒲⒩ ⒡⒪⒳ ⒥⒰⒨⒫⒮ ⒪⒱⒠⒭ ⒯⒣⒠ ⒧⒜⒵⒴ ⒟⒪⒢",
}

// SillyString replaces a value with a random string that is known to potentially cause problems.
// Example: Hello -> `⁄€‹›ﬁﬂ‡°·‚—±
type SillyString struct{}

// ID returns mutator's 3-digit ID.
func (m *SillyString) ID() string {
	return "019"
}

// Name returns the mutator's name.
func (m *SillyString) Name() string {
	return "Silly String"
}

// Description returns the mutator's description.
func (m *SillyString) Description() string {
	return "Replace value with random silly string"
}

// CompatibleKinds returns the data types that the mutator can mutate.
func (m *SillyString) CompatibleKinds() []reflect.Kind {
	return []reflect.Kind{reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Invalid}
}

// Mutate returns a random string that is known to potentially cause problems.
// Example: `⁄€‹›ﬁﬂ‡°·‚—±
func (m *SillyString) Mutate(subject reflect.Value, r *rand.Rand) interface{} {
	return sillyStringPayloads[r.Intn(len(sillyStringPayloads))]
}
