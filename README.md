# BlogWebsite
create website with golang and reactjs css and mysql

when want to use module from another folder
ex. in folder A use mod folder B

$cd folder A

$go mod edit -replace=modOfB/floder=../floder

ex. $ go mod edit -replace=blogwebsite/common=../common

$go mod tidy

finish....
