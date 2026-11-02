function voteing(age){
    let result = confirm("Do you want print?")

    if (age>=18 && result){
        console.log("Eligible for vote")
    }
}
voteing(20)
