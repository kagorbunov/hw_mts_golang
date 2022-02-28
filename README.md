data:size:name>>>><<<<hash

main [enc, dec] --sourceFile=name --hashFile=sign --saveTo=out.file
HomeWork 
-create ValidateSign() -> lesson7/mycrypt/Decryptor.py //for check file with signature 
-create NewFileSourceFromSignatureSha256(fileSign) -> lesson7/signature/signatureSha256.go //for parse fail with hash string and next build signature
-refactor Equals(ss contract.Signature) -> lesson7/signature/signatureSha256.go //for compare signature, but with out date

lesson7.exe --source-file=source.txt --hash-file=hash.txt --out-file=sign.txt dec
