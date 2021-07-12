## Поиск префикса/суфикса в строке из списка

### Описание:
Если есть в слове префикс/суфикс из массива - он возвращает это слово + остаток 2м параметром

### Примеры:
С учетом регистра
```
arr := []string{"СЛОВО","АВТО"}

Prefix("BMWАВТО", arr) => "", "", "no match found"
Prefix("АВТОBMW", arr) => "АВТО", "BMW", nil
Prefix("АвтоBMW", arr) => "", "", "no match found"

Suffix("BMWАВТО", arr) => "АВТО", "BMW", nil
Suffix("BMWАвто", arr) => "", "", "no match found"
Suffix("АВТОBMW", arr) => "", "", "no match found"
```
Без учета регистра
```
arr := []string{"СЛОВО","АВТО", ""}

PrefixWithoutCase("BMWАВТО", arr) => "", "", "no match found"
PrefixWithoutCase("АВТОBMW", arr) => "АВТО", "BMW", nil
PrefixWithoutCase("АвтоBMW", arr) => "Авто", "BMW", nil

SuffixWithoutCase("BMWАВТО", arr) => "АВТО", "BMW", nil
SuffixWithoutCase("BMWАвто", arr) => "Авто", "BMW", nil
SuffixWithoutCase("АВТОBMW", arr) => "", "", "no match found"
```
В массиве пустая строка - ИГНОРИРУЕТСЯ
