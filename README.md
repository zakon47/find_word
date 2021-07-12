## Поиск префикса/суфикса в строке из списка

### Описание:
Если есть в слове префикс/суфикс из массива - он возвращает это слово + остаток 2м параметром

### Примеры:
С учетом регистра
```
arr := []string{"СЛОВО","АВТО"}

HasPrefix("BMWАВТО", arr) => "", "", "no match found"
HasPrefix("АВТОBMW", arr) => "АВТО", "BMW", nil
HasPrefix("АвтоBMW", arr) => "", "", "no match found"

HasSuffix("BMWАВТО", arr) => "АВТО", "BMW", nil
HasSuffix("BMWАвто", arr) => "", "", "no match found"
HasSuffix("АВТОBMW", arr) => "", "", "no match found"
```
Без учета регистра
```
arr := []string{"СЛОВО","АВТО", ""}

HasPrefixWithoutCase("BMWАВТО", arr) => "", "", "no match found"
HasPrefixWithoutCase("АВТОBMW", arr) => "АВТО", "BMW", nil
HasPrefixWithoutCase("АвтоBMW", arr) => "Авто", "BMW", nil

HasSuffixWithoutCase("BMWАВТО", arr) => "АВТО", "BMW", nil
HasSuffixWithoutCase("BMWАвто", arr) => "Авто", "BMW", nil
HasSuffixWithoutCase("АВТОBMW", arr) => "", "", "no match found"
```
В массиве пустая строка - ИГНОРИРУЕТСЯ
