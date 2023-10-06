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

const userTotals = {};

transactions.forEach(transaction => {
  const { userId, total } = transaction;
  if (userTotals[userId]) {
    userTotals[userId] += total;
  } else {
    userTotals[userId] = total;
  }
  console.log(userTotals)
});

// Map the users to include the total from the dictionary
const result = users.map(user => ({
  id: user.id,
  name: user.name,
  total: userTotals[user.id] || 0, // Default to 0 if no transactions for the user
}));

console.log(result);




     