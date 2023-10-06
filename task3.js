const printSquare = (num) => {
    for (let i = 0; i<num; i++){
        let pattern = ''
        for (let j=0; j<num; j++) {
            pattern+='*'
        }
        console.log(pattern)
    }
}

printSquare(5)