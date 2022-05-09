let examplePokemonDict = {
    1: {
        name: "Bulbasaur",
        number: 1,
        amount: "$10,800",
        due: "12/05/1995",
    },
    2: {
        name: "Ivysaur",
        number: 2,
        amount: "$8,000",
        due: "10/31/2000",
    },
    3: {
        name: "Venusaur",
        number: 3,
        amount: "$9,500",
        due: "07/22/2003",
    },
    4: {
        name: "Charmander",
        number: 4,
        amount: "$14,000",
        due: "09/01/1997",
    },
    5: {
        name: "Charmeleon",
        number: 5,
        amount: "$4,600",
        due: "01/27/1998",
    },
    6: {
        name: "Charizard",
        number: 6,
        amount: "$4,600",
        due: "01/27/1998",
    }
  }
  export function getAllExamplePokemonDict(){
    return examplePokemonDict
  }
  export function getAExamplePokemonDict(number){
    return examplePokemonDict[number]
  }
  export function printAExamplePokemonDict(number){
    console.log(examplePokemonDict[number]);
  }
  printAExamplePokemonDict(1);