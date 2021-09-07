# Todo

For Windows Machin
Change this:

```go

func NewScreenPrinter() *ScreenPrinter {
	w := new(tabwriter.Writer)
	w.Init(color.Output, 0, 8, 1, ' ', 0) //Changed from os.Stdout to color.Output when compiled on Windows.
	//w.Init(color.Output, 5, 0, 1, ' ', tabwriter.StripEscape)
```

TO:

```go

func NewScreenPrinter() *ScreenPrinter {
	w := new(tabwriter.Writer)
	// w.Init(color.Output, 0, 8, 1, ' ', 0) //Changed from os.Stdout to color.Output when compiled on Windows.
	w.Init(color.Output, 5, 0, 1, ' ', tabwriter.StripEscape)
```

### Shell CMD:

```sh

go build .

./todo add -s "Come to Cafe One" -d "Come CafeOne by 8:00am"

./todo list

./todo cleanup

./todo done -a={{task id}}

./todo undone -a={{task id}}

```
