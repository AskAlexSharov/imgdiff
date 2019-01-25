`imgdiff` - cli tool to calculate difference between images.
Results are stable for next transformations: "scale-up/scale-down", color adjustment
Returns % of difference

## Install

```bash
go install -u github.com/AskAlexSharov/imgdiff
```

## Use

```bash
imgdiff ./test-png-original.png ./test-png-damaged.png
Difference: 3%
```

Or you can use URL of image
```bash
imgdiff https://urloffile.com/file1.jpg https://urloffile.com/file2.jpg
Difference: 0% 
```
## Features

- JPEG
- PNG 
- Local files and download from Url