function main(ctx, args) {
    console.log(JSON.stringify(ctx))
    console.log("==========")
    console.log(JSON.stringify(args))
    let name = args.name || 'stranger'
    let greeting = 'Hello Two ' + name + '!'
    console.log(greeting)
    return {"body": greeting}
  }

exports.main = main
