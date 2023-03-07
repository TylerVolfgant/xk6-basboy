xk6-basboy

This is a k6 extension using the xk6 system.
exclamation This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK!

About Counters:

This projects implements a singular local CounterGlobal() that goes up from 1,

counterPRGS() from startValue and increases according to the increment,

counterFormat() from startValue and increases according to the increment, you can set maxValue, and set format like "0000".


It will return the current value before increasing it, and each VU will get a different value. Which means that it can be used to iterate over an array, where:

    only one element will be used by each VU
    the array doesn't need to be sharded between the VUs it will "dynamically" balance between them even if some elements take longer to process

This totally doesn't work in distributed manner, so if there are multiple k6 instances you will need a separate service (API endpoint) which to do it.

Predominantly because of the above this is very unlikely to ever get in k6 in it's current form, so please don't open issues :D.

About random values:


In this project you can find several methods:

randomItem() - return random item from array;

randomIntBetween() - return random int between range;

uuidv4() - return random uuid v4;

alphanumeric() - random Alphanumeric string with length, and bool to uppercase;

alphabetic() - random Alphabetic string with length, and bool to uppercase;

numeric() -  random Numeric string with length;

alphanumericAndSymbolic() - random AlphanumericAndSymbolic string with length;

hexadecimal() - random Hexadecimal string with length; 

randomString() - random RandomString string with length, and bool to uppercase;


About date time:


rnow() - function to call current date time, and you can set some timezones, or use offset to set date time with '+' or '-' unit(secs, mins, hours, days, weeks, months, years);

You can also set an output template for date time.

Build

To build a k6 binary with this extension, first ensure you have the prerequisites:

    gvm
    Git

Then, install xk6 and build your custom k6 binary with the Kafka extension:

    Install xk6:

$ go install go.k6.io/xk6/cmd/xk6@latest

    Build the binary:

$ xk6 build --with github.com/TylerVolfgant/xk6-basboy@latest 

example

import basboy from 'k6/x/basboy'


export let options = {
vus: 1,
iterations: 3,
}

let counterPRGS1New = basboy.counterPRGS(1,10);

let counterFormatNew = basboy.counterFormat(10, 5, 1000, "0000");



export default function () {



let arrOp = ["tree", "one", "four", "666", "pantera"];


//counters
console.log("counterGlobal !!! ==== ", basboy.counterGlobal()); //1

console.log("counterPRGS !!! ==== ", counterPRGS1New()); //2

console.log("counterFormat !!! ==== ", counterFormatNew()); //3



//random values
console.log("randomItem ==== ", basboy.randomItem(arrOp)); //4

console.log("randomIntBetween ==== ", basboy.randomIntBetween(1, 89)); //5

console.log("uuidv4 ==== ", basboy.uuidv4()); //6

console.log("alphanumeric ==== ", basboy.alphanumeric(5)); //7

console.log("alphanumeric upper ==== ", basboy.alphanumeric(5, true)); //7

console.log("alphabetic ==== ", basboy.alphabetic(9)); //8

console.log("alphabetic upper ==== ", basboy.alphabetic(9, true)); //8

console.log("numeric ==== ", basboy.numeric(6)); //9

console.log("alphanumeric_and_symbolic up ==== ", basboy.alphanumericAndSymbolic(7)); //10

console.log("hexadecimal ==== ", basboy.hexadecimal(10)); //11

console.log("randomString ==== ", basboy.randomString(12)); //12

console.log("randomString ==== ", basboy.randomString(12, true)); //12



//time format, offset, timezone //13

console.log("now default ==== ", basboy.rnow());

console.log("now offset 1 days  ==== ", basboy.rnow({"offset": "+1days" }));

console.log("now offset 3 days  ==== ", basboy.rnow({"offset": "-3days" }));

console.log("now offset +1h ==== ", basboy.rnow({"offset": "+1hours" }));

console.log("now offset -1h ==== ", basboy.rnow({"offset": "-1hours" }));

console.log("now offset +5m ==== ", basboy.rnow({"offset": "+5mins" }));

console.log("now offset -5m ==== ", basboy.rnow({"offset": "-5mins" }));

console.log("now offset +20sec ==== ", basboy.rnow({"offset": "+20secs" }));

console.log("now offset -20sec ==== ", basboy.rnow({"offset": "-20secs" }));

console.log("now offset +1 years  ==== ", basboy.rnow({"format": "YYYY-MM-DD hh:mm:ss", "timezone":"Africa/Accra", "offset": "+1years" }));

console.log("now offset -10 years  ==== ", basboy.rnow({"format": "YYYY-MM-DD hh:mm:ss", "timezone":"UTC", "offset": "-10years" }));

console.log("now kitchen  ==== ", basboy.rnow({"format": "h:mmam/pm"}));

console.log("now -1 days timezone: EST ==== ", basboy.rnow({"offset": "-1days", "timezone": "EST", "format": "yyyy-MM-DD hh:mm:ssZ" }));


}
