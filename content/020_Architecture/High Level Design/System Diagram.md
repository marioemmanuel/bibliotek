---
# Header 1
## Header 2
### Header 3
#### Header 4
##### Header 5

Lorem Ipsum

---

This is a table


<center>
| Item              | In Stock | Price |
| :---------------- | :------: | ----: |
| Python Hat        |   True   | 23.99 |
| SQL Hat           |   True   | 23.99 |
| Codecademy Tee    |  False   | 19.99 |
| Codecademy Hoodie |  False   | 42.99 |
</center>


This is source code 


```go
func init() {
	// Define the command line flag
	// The flag is -path, with a default value and a short description
	flag.StringVar(&rootPath, "path", "", "Path to the folder to be exposed")

	// Parse the flags
	flag.Parse()

	// Check if the path is provided
	if rootPath == "" {
		fmt.Println("You must provide a file root path using the -path flag.")
		os.Exit(1)
	}
}
```

Text Centered to the right 

<p style="text-align: right;">THIS TEXT GOES RIGHT</p>

