# whut
Command line word lookup

Currently looks up English word on http://glosbe.com and lists its Russian translations.

Configs, more features, and precompiled binaries pending.

## Usage
```
whut <word to translate>
```

## Example
```
whut shortbread
```
Output:
```
1. песочное печенье
2. very rich thick butter cookie
3. A type of biscuit (cookie), popular in Britain, traditionally made from one part sugar, two parts butter and three parts flour
```
## Build
```
cd whut
go get jaytaylor.com/html2text
go build -ldflags "-w"
```
## Unit Tests
```
go test
```
