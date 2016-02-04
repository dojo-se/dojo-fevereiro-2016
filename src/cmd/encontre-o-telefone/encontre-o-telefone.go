package main

func getTelefone(strTelefone string) (numTelefone string) {
    numTelefone = ""
    dict := make (map[string] string)
    dict["A"] = "2"
    dict["B"] = "2"
    dict["C"] = "2"
    dict["D"] = "3"
    dict["-"] = "-"
    dict[" "] = ""
    
    for index := 0; index < len(strTelefone); index++ {
        str, ok := dict[strTelefone[index:index+1]]
        
        if !ok {
            return "test"     
        }
        numTelefone += str
    }
    
    return
}
