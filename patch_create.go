package main

import (
    "io/ioutil"
    "strings"
)

func main() {
    content, _ := ioutil.ReadFile("/Users/andregb/Documents/WORK/nai/nei-api/internal/handlers/catalog.go")
    strContent := string(content)
    strContent = strings.ReplaceAll(strContent, "database.DB.Create(&neu).Error", "func() error {\n\t\tif neu.MarcaID != nil && *neu.MarcaID == 0 {\n\t\t\tneu.MarcaID = nil\n\t\t}\n\t\treturn database.DB.Create(&neu).Error\n\t}()")
    
    // Also fix UpdateNeumatico
    strContent = strings.ReplaceAll(strContent, "database.DB.Save(&neu).Error", "func() error {\n\t\tif neu.MarcaID != nil && *neu.MarcaID == 0 {\n\t\t\tneu.MarcaID = nil\n\t\t}\n\t\treturn database.DB.Save(&neu).Error\n\t}()")

    ioutil.WriteFile("/Users/andregb/Documents/WORK/nai/nei-api/internal/handlers/catalog.go", []byte(strContent), 0644)
}
