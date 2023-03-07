import basboy from 'k6/x/basboy'

export let options = {
    vus: 1,
    iterations: 5,
}

let counterPRGSNew = basboy.counterPRGS(1,10);
let counterFormatNew = basboy.counterFormat(10, 10, 40, "00");

export default function () {
let arrOp = ["tree", "one", "four", "666", "pantera"];
////counters
console.log("counterGlobal !!! ==== ", basboy.counterGlobal()); //1
console.log("counterPRGS !!! ==== ", counterPRGSNew()); //2
console.log("counterFormat !!! ==== ", counterFormatNew()); //3

////random values
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

////time format //13
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
