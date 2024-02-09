function main(args) {
    console.log(JSON.stringify(args))
    let name = args.name || 'stranger'
    let greeting = 'Hello Two ' + name + '!'
    console.log(greeting)
    return {"body": greeting}
  }

exports.main = main
