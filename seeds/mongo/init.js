db = new Mongo().getDB('fakedata');

db.createCollection('users', { capped: false });

db.users.insert([
    {
        username: 'leonardo',
        password: '12345',
        country: 'Venezuela',
        age: '25',
        role: 'Technical Leader | Full Stack Developer',
        hobby: 'Bicicle, Run, Volley, Soccer',
    },
    {
        username: 'pepe',
        password: '67890',
        country: 'United States',
        age: '18',
        role: 'Full Stack Developer',
        hobby: 'Run, GYM, Football',
    },
    {
        username: 'John',
        password: '67890',
        country: 'United States',
        age: '33',
        role: 'Killer',
        hobby: 'Sing, GYM, Racing',
    },
    {
        username: 'Wick',
        password: '67890',
        country: 'United States',
        age: '40',
        role: 'Scrum Master',
        hobby: 'Football',
    },
]);
