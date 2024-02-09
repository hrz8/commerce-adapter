function main(args) {
    let name = args.name || 'stranger'
    let greeting = 'Hello One ' + name + '!'
    console.log(greeting)
    return {"body": greeting}
  }

exports.main = main
