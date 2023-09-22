const transactions = [
    {
      id: 1,
      userId: 1,
      total: 15000,
      date: '2023-01-01'
    }, {
      id: 2,
      userId: 2,
      total: 20000,
      date: '2023-01-02'
    }, {
        id: 3,
        userId: 1,
        total: 10000,
        date: '2023-01-03'
      }, {
        id: 4,
        userId: 1,
        total: 5000,
        date: '2023-01-02'
      }, {
        id: 5,
        userId: 2,
        total: 5000,
        date: '2023-01-05'
      }, {
        id: 6,
        userId: 3,
        total: 75000,
        date: '2023-01-04'
      }
     ]
const users = [
{
    id: 1,
    name: "Udin"
}, {
    id: 2,
    name: "Richard"
}, {
    id: 3,
    name: "Tono"
}
]

const result = [{}]

for (const value of Object.values(transactions)){
    result.push(value.total)
}


     