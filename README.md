# rfc2822

## Usage
```
func main() {
	headers := map[string]string{
		"From":         "Sample",
		"Subject":      "Test",
		"Content-Type": "text/html; charset=UTF-8",
	}
	body := "<html><body>test</body><html>"
	m := rfc2822.Message{Headers: headers, Body: body}

	fmt.Println(m.Create())
}
```
```
From: Sample
Subject: Test
Content-Type: text/html; charset=UTF-8

<html><body>test</body><html>
```

Use no-ASCII header
```
func main() {
	headers := map[string]string{
		"From":         "Sample",
		"Subject":      rfc2822.EncodeRFC2047(rfc2822.UTF8, []byte("テスト")),
		"Content-Type": "text/html; charset=UTF-8",
	}
	body := "<html><body>test</body><html>"
	m := rfc2822.Message{Headers: headers, Body: body}

	fmt.Println(m.Create())
}
```
```
From: Sample
Subject: =?UTF-8?B?44OG44K544OI?=
Content-Type: text/html; charset=UTF-8

<html><body>test</body><html>
```
