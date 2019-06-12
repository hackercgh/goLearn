package main

import (
	"crypto/des"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"lib/publib/github.com/wonderivan/logger"
	"os"
)

/*var key = []byte{
0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,
0x08,0x09,0x1A,0x1B,0x1C,0x1D,0x1E,0x1F,
0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,
0x08,0x09,0x1A,0x1B,0x1C,0x1D,0x1E,0x1F }*/

func main() {
	tripleDESKey := []byte("111111111111111111111111")
	val := []byte("0123456789ABCDEF0123456789ABCDEF")
	var dst []byte
	var dst1 []byte
	tmp := make([]byte, 8)
	tmp1 := make([]byte, 8)
	/*cb,err := des.NewTripleDESCipher(key)
	if err != nil {
		logger.Error("des.NewTripleDESCipher failed [%v]",err)
		os.Exit(1)
	}
	cb.Encrypt(dst,val)
	logger.Debug("dst=[%v]",dst)
	ede2Key := []byte("example key 1234")
	*/
	/*ede2Key := []byte("example key 1234")
	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)*/

	cb, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}
	l := len(val)
	for i := 0; i < l; {
		copy(tmp, val[i:i+8])
		logger.Debug("src[%s]", tmp)
		cb.Encrypt(tmp1, tmp)
		i += 8
		dst = append(dst, tmp1...)
	}

	logger.Debug("dst=%s", dst)
	for i := 0; i < l; {
		copy(tmp, dst[i:i+8])
		cb.Decrypt(tmp1, tmp)
		dst1 = append(dst1, tmp1...)
		i += 8
	}

	logger.Debug("dst1=[%s]", dst1)
	dst1 = (dst1)[0:0]
	logger.Debug("dst1=[%s]", dst1)

	//src := []byte("Hello Gopher!")
	src := []byte("0123456789ABCDEF0123456789ABCDEF")

	d := make([]byte, hex.EncodedLen(len(src)))
	//var s []byte
	s := make ([]byte,32)
	hex.Encode(d, src)
	hex.Decode(s,d)
	logger.Debug("%s", d)
	logger.Debug("%s",s)

	h := make ([]byte,100)
	s1 := make ([]byte,100)
	hex.Encode(h,src)
	logger.Debug("h=[%s]",h)
	hex.Decode(s1,h)
	logger.Debug("s=[%s]",s1)

	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"Person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age,attr"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}
	v :=Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Height = 32.31
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	//output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	logger.Debug("\nage=%v",v.Age)
	//os.Exit(1)
/*	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}*/

	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v1 := Result{Name: "none", Phone: "none"}
	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<AAA>aaaaaaaaaaaaa</AAA>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err = xml.Unmarshal([]byte(data), &v1)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v1.Name)
	fmt.Printf("Phone: %q\n", v1.Phone)
	fmt.Printf("Email: %v\n", v1.Email)
	fmt.Printf("Groups: %v\n", v1.Groups)
	fmt.Printf("Address: %v\n", v1.Address)
}
