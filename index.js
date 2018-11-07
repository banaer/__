class FooPromise extends Promise {
    constructor(executor) {
        super((res,rej)=>{
            console.log('super')
            if(executor){
                return executor(res,rej)
            }
            setTimeout(()=>{
                res('111')
            },1000)

        });
    }
    then(res){

        return super.then(res)

    }
}

var fooPromise = new FooPromise().then(d=>{
    console.log(d)
    return 1;
}).then(m=>{
    console.log('m',m)
});

console.log(fooPromise instanceof FooPromise)
