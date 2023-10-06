const printSegitiga = (num) =>{
    for (let i=0; i<=num;i++){
        let pattern = ''
        for (let j=1; j<=i; j++){
            pattern+=j+''
        }
        console.log(pattern)
    }
}

printSegitiga(5)