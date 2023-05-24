import axios from 'axios';
import config from './config.js';

const data = []

async function benchMe () {
    const time1 = performance.now();
    const response = await axios.get(config.url)
    await response.data
    return performance.now() - time1;
}

for (let i = 0; i < 100; i++) {
    const time = await benchMe()
    data.push(time)
}
   
console.log(data.reduce((a, b) => a + b, 0) / data.length)