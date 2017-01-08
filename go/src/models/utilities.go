package models

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"io"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
	CURRENT_VERSION   = "0.0.0"
)

type StringInterface map[string]interface{}
type StringMap map[string]string
type StringArray []string
type EncryptStringMap map[string]string

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

func NewRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}

// GetMillis is a convience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MapToJson converts a map to a json string
func MapToJson(objmap map[string]string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

// MapFromJson will decode the key/value pair map
func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}

func ArrayToJson(objmap []string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ArrayFromJson(data io.Reader) []string {
	decoder := json.NewDecoder(data)

	var objmap []string
	if err := decoder.Decode(&objmap); err != nil {
		return make([]string, 0)
	} else {
		return objmap
	}
}

func ArrayFromInterface(data interface{}) []string {
	stringArray := []string{}

	dataArray, ok := data.([]interface{})
	if !ok {
		return stringArray
	}

	for _, v := range dataArray {
		if str, ok := v.(string); ok {
			stringArray = append(stringArray, str)
		}
	}

	return stringArray
}

func StringInArray(a string, array []string) bool {
	for _, b := range array {
		if b == a {
			return true
		}
	}
	return false
}

func StringInterfaceToJson(objmap map[string]interface{}) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func StringInterfaceFromJson(data io.Reader) map[string]interface{} {
	decoder := json.NewDecoder(data)

	var objmap map[string]interface{}
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]interface{})
	} else {
		return objmap
	}
}

func StringToJson(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func StringFromJson(data io.Reader) string {
	decoder := json.NewDecoder(data)

	var s string
	if err := decoder.Decode(&s); err != nil {
		return ""
	} else {
		return s
	}
}

func IsLower(s string) bool {
	return strings.ToLower(s) == s
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil && IsLower(email)
}

func IsValidDomain(domain string) bool {
	return IsLower(domain) && IsValidAlphaNum(domain, true)
}

var reservedName = []string{
	"signup",
	"login",
	"admin",
	"channel",
	"post",
	"api",
	"oauth",
}

func IsValidChannelIdentifier(s string) bool {
	return IsValidAlphaNum(s, true)
}

func IsValidOrganisationIdentifier(s string) bool {

	return IsValidAlphaNum(s, true)
}

var validAlphaNumUnderscore = regexp.MustCompile(`^[a-z0-9]+([a-z\-\_0-9]+|(__)?)[a-z0-9]+$`)
var validAlphaNum = regexp.MustCompile(`^[a-z0-9]+([a-z\-0-9]+|(__)?)[a-z0-9]+$`)

func IsValidAlphaNum(s string, allowUnderscores bool) bool {
	var match bool
	if allowUnderscores {
		match = validAlphaNumUnderscore.MatchString(s)
	} else {
		match = validAlphaNum.MatchString(s)
	}

	return match
}

func Etag(parts ...interface{}) string {

	Etag := CURRENT_VERSION

	for _, part := range parts {
		Etag += fmt.Sprintf(".%v", part)
	}

	return Etag
}

var validHashtag = regexp.MustCompile(`^(#\pL[\pL\d\-_.]*[\pL\d])$`)
var puncStart = regexp.MustCompile(`^[^\pL\d\s#]+`)
var hashtagStart = regexp.MustCompile(`^#{2,}`)
var puncEnd = regexp.MustCompile(`[^\pL\d\s]+$`)

func ParseHashtags(text string) (string, string) {
	words := strings.Fields(text)

	hashtagString := ""
	plainString := ""
	for _, word := range words {
		// trim off surrounding punctuation
		word = puncStart.ReplaceAllString(word, "")
		word = puncEnd.ReplaceAllString(word, "")

		// and remove extra pound #s
		word = hashtagStart.ReplaceAllString(word, "#")

		if validHashtag.MatchString(word) {
			hashtagString += " " + word
		} else {
			plainString += " " + word
		}
	}

	if len(hashtagString) > 1000 {
		hashtagString = hashtagString[:999]
		lastSpace := strings.LastIndex(hashtagString, " ")
		if lastSpace > -1 {
			hashtagString = hashtagString[:lastSpace]
		} else {
			hashtagString = ""
		}
	}

	return strings.TrimSpace(hashtagString), strings.TrimSpace(plainString)
}

// func IsFileExtImage(ext string) bool {
// 	ext = strings.ToLower(ext)
// 	for _, imgExt := range IMAGE_EXTENSIONS {
// 		if ext == imgExt {
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetImageMimeType(ext string) string {
// 	ext = strings.ToLower(ext)
// 	if len(IMAGE_MIME_TYPES[ext]) == 0 {
// 		return "image"
// 	} else {
// 		return IMAGE_MIME_TYPES[ext]
// 	}
// }

func ClearMentionTags(post string) string {
	post = strings.Replace(post, "<mention>", "", -1)
	post = strings.Replace(post, "</mention>", "", -1)
	return post
}

var UrlRegex = regexp.MustCompile(`^((?:[a-z]+:\/\/)?(?:(?:[a-z0-9\-]+\.)+(?:[a-z]{2}|aero|arpa|biz|com|coop|edu|gov|info|int|jobs|mil|museum|name|nato|net|org|pro|travel|local|internal))(:[0-9]{1,5})?(?:\/[a-z0-9_\-\.~]+)*(\/([a-z0-9_\-\.]*)(?:\?[a-z0-9+_~\-\.%=&amp;]*)?)?(?:#[a-zA-Z0-9!$&'()*+.=-_~:@/?]*)?)(?:\s+|$)$`)
var PartialUrlRegex = regexp.MustCompile(`/([A-Za-z0-9]{26})/([A-Za-z0-9]{26})/((?:[A-Za-z0-9]{26})?.+(?:\.[A-Za-z0-9]{3,})?)`)

var SplitRunes = map[rune]bool{',': true, ' ': true, '.': true, '!': true, '?': true, ':': true, ';': true, '\n': true, '<': true, '>': true, '(': true, ')': true, '{': true, '}': true, '[': true, ']': true, '+': true, '/': true, '\\': true}

func IsValidHttpUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "http://") != 0 && strings.Index(rawUrl, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}

func IsValidHttpsUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}

func IsValidTurnOrStunServer(rawUri string) bool {
	if strings.Index(rawUri, "turn:") != 0 && strings.Index(rawUri, "stun:") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUri); err != nil {
		return false
	}

	return true
}

func IsSafeLink(link *string) bool {
	if link != nil {
		if IsValidHttpUrl(*link) {
			return true
		} else if strings.HasPrefix(*link, "/") {
			return true
		} else {
			return false
		}
	}

	return true
}

func IsValidWebsocketUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "ws://") != 0 && strings.Index(rawUrl, "wss://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}
