`imgdiff` - cli tool to calculate difference between images.
Results are stable for next transformations: "scale-up/scale-down", color adjustment, quality reduce
Returns % of difference

Thanks https://github.com/Nr90/imgsim/ for Average and Difference hash calculations 

### Install

```bash
go get -u github.com/AskAlexSharov/imgdiff
```

### Use

```bash
imgdiff ./avg/test-png-original.png ./avg/test-png-damaged.png
Difference: 2%
```

Or you can use URL of image
```bash
imgdiff https://raw.githubusercontent.com/AskAlexSharov/imgdiff/master/avg/test-png-original.png https://raw.githubusercontent.com/AskAlexSharov/imgdiff/master/avg/test-png-damaged.png
Difference: 2% 
```
### Features

- JPEG
- PNG 
- Local files and download from Url

### License

`imgdiff` is [MIT licensed](./LICENSE)