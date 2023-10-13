const students = [
    {
      id: 1,
      name: "Heru",
      score: 90
    }, {
      id: 2,
      name: "Herman",
      score: 80
    }, {
      id: 3,
      name: "Budi",
      score: 85
    }, {
      id: 4,
      name: "Tika",
      score: 87
    }, {
      id: 5,
      name: "Junaedi",
      score: 79
    }, {
      id: 6,
      name: "Ahmad",
      score: 82
    }, {
      id: 7,
      name: "Dani",
      score: 81
    }
   ]
   
   for (let i = 0; i < students.length - 1; i++) {
     for (let j = 0; j < students.length - i - 1; j++) {
       if (students[j].score > students[j + 1].score) {
         const temp = students[j];
         students[j] = students[j + 1];
         students[j + 1] = temp;
       }
     }
   }
   console.log(students)