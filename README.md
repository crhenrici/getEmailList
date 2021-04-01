# getEmailList

Small tool to get a list of mails from a given list with multiple employees, duplicates included

### Details

With the latest update I made use of the [excelize](https://github.com/360EntSecGroup-Skylar/excelize) library. Therefore I can directly access the column needed. To get all the values of the column I simply iterate through the rows.
The GUI is made of a button, a textfield and a menu to open a file dialog. The design is kept simple. The GUI is made with [winc](https://github.com/tadvi/winc) library.

### TODOS
- Encode user inputs with umlauts to UTF-8 to replace umlauts on case to case basis[DONE] (Due to the use of excel sheets, all the strings are UTF-8 encoded)
- Improve programm, e.g. use column with the needed names instead of a hardcoded column [DONE]
