const fs = require("fs");

async function main() {
    const writeStream = fs.createWriteStream("small.csv");
    writeStream.write("index, data\n")
    for (let i = 0; i < 1e6; i++){
        const writeable = writeStream.write(`${i}, ${i * 2}\n`)

        if (!writeable) {
            await new Promise((resolve) =>
                writeStream.once("drain", resolve)
            );
        }
    }
    writeStream.end();
}

main().catch((err) => {
    console.error(err)
    process.exit(1)
})