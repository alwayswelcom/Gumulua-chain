
// file system module to perform file operations
const fs = require('fs');
const R = require('ramda');
const util = require('util')
const atob = require('atob')
const btoa = require('btoa')

var genesisold = require('./genesis.old.json');
var genesisnew = require('./genesis.new.json');

console.log('old:', genesisold.app_state)
console.log('new:', genesisnew.app_state)

// TODO double ids for owned cards
// TODO BUY CARD SCHEMES OVERWRITE REAL CARDS
genesisnew.app_state.auth.accounts = genesisold.app_state.auth.accounts
genesisnew.app_state.cardservice.users = genesisold.app_state.cardservice.users
genesisnew.app_state.cardservice.addresses = genesisold.app_state.cardservice.addresses
genesisnew.app_state.cardservice.card_records = genesisold.app_state.cardservice.card_records
/*
genesisnew.app_state.cardservice.users.push(JSON.parse(JSON.stringify(genesisnew.app_state.cardservice.users[0])))
*/
//console.log('merged:', genesisnew.app_state)
console.log('users before sanitation:', genesisnew.app_state.cardservice.users)

genesisnew.app_state.cardservice.users = R.map(function(x) { x.OwnedCards = x.OwnedCards ? R.uniq(x.OwnedCards) : []; return x }, genesisnew.app_state.cardservice.users)
genesisnew.app_state.cardservice.users = R.map(function(x) { x.OwnedCardSchemes = x.OwnedCardSchemes ? R.uniq(x.OwnedCardSchemes) : []; return x }, genesisnew.app_state.cardservice.users)
genesisnew.app_state.cardservice.users = R.map(function(x) { x.OwnedCardSchemes = x.OwnedCardSchemes && x.OwnedCards ? R.without(x.OwnedCards, x.OwnedCardSchemes) : x.OwnedCardSchemes; return x }, genesisnew.app_state.cardservice.users)

// card model merger
let id = 0

genesisnew.app_state.cardservice.card_records = R.map(x => {
  //console.log('decoded:', atob(x.Content != null ? x.Content : btoa('{}')))
  console.log(id)
  id++

  let content = JSON.parse(atob(x.Content != null ? x.Content : btoa('{}')))

  //console.log(content)

  if (content.Action) {
    R.filter(ability => {
      let checkstring = JSON.stringify(ability)
      if (checkstring.includes("Arm") ||
          checkstring.includes("Harm") ||
          checkstring.includes("Repair") ||
          checkstring.includes("Kill") ||
          checkstring.includes("Heal")  ) {
          if (!checkstring.includes("Target")) {
            console.log("FAIILLLL:", ability)
            //console.log(util.inspect(content.Action.Effects, {showHidden: false, depth: null}))
            return {}
          }
      }
      return ability
    }, content.Action.Effects)
    //content.Action.RulesTexts = []
    //content.Action.Effects = []
    //console.log(util.inspect(content.Action.Effects, {showHidden: false, depth: null}))
  }
  else if (content.Place && content.Place.Abilities) {
    R.filter(ability => {
      let checkstring = JSON.stringify(ability)

      if (checkstring.includes("Arm") ||
          checkstring.includes("Harm") ||
          checkstring.includes("Repair") ||
          checkstring.includes("Kill") ||
          checkstring.includes("Heal") ) {

        if (!checkstring.includes("Target") ) {
          console.log("FAIILLLL:", ability)
          return {}
          //console.log(util.inspect(content.Place.Abilities, {showHidden: false, depth: null}))
        }

      }
      return ability
    }, content.Place.Abilities)
    //content.Place.RulesTexts = []
    //content.Place.Abilities = []
    //console.log(util.inspect(content.Place.Abilities, {showHidden: false, depth: null}))
  }
  else if (content.Headquarter) {
    R.filter(ability => {
      checkstring = JSON.stringify(ability)

      if (checkstring.includes("Arm")||
          checkstring.includes("Harm")||
          checkstring.includes("Repair") ||
          checkstring.includes("Kill") ||
          checkstring.includes("Heal")  ) {
        if (!checkstring.includes("TARGET") ) {
          console.log("FAIILLLL:", ability)
          return {}
          //console.log(util.inspect(content.Headquarter.Abilities, {showHidden: false, depth: null}))
        }

      }
      return ability
    }, content.Headquarter.Abilities)
    //content.Headquarter.RulesTexts = []
    //content.Headquarter.Abilities = []
    //console.log(util.inspect(content.Headquarter.Abilities, {showHidden: false, depth: null}))
  }
  else if (content.Entity) {
    R.filter(ability => {
      checkstring = JSON.stringify(ability)
      if (checkstring.includes("Arm") ||
          checkstring.includes("Harm") ||
          checkstring.includes("Repair") ||
          checkstring.includes("Kill") ||
          checkstring.includes("Heal") ) {
        if (!checkstring.includes("Target") ) {
          console.log("FAIILLLL:", ability)
          return {}
          console.log(util.inspect(content.Entity.Abilities, {showHidden: false, depth: null}))
        }

      }
      return ability
    }, content.Entity.Abilities)
    //content.Entity.RulesTexts = []
    //content.Entity.Abilities = []
    //console.log(util.inspect(content.Entity.Abilities, {showHidden: false, depth: null}))
  }

  x.Content = btoa(JSON.stringify(content))
  return x
}, genesisnew.app_state.cardservice.card_records)

/*
console.log('users written in the file:', genesisnew.app_state.cardservice.users)
//console.log('addresses:', genesisnew.app_state.cardservice.addresses)
console.log('cards:', R.map(x => {
    return atob(x.Content ? x.Content : '')
  }, genesisnew.app_state.cardservice.card_records))
*/
//ids = R.map(x => x.Owner, genesisnew.app_state.cardservice.card_records)
//console.log(ids)

fs.writeFileSync('./genesis.merged.json', JSON.stringify(genesisnew) , 'utf-8');
