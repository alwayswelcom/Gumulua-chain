#!/bin/bash

printf 'Y\nasdfasdf\n' | cscli tx cardservice save-card-content 1 '{"Entity":{"CardName":"Archer","Tags":["RANGE"],"Text":"Ping","CostType":{"Food":true,"Lumber":true},"CastingCost":2,"Abilities":[{"ActivatedAbility":{"AbilityCost":1,"MultipleUse":true,"Effects":[{"TargetEffect":{"EntityTargetEffect":{"EntitySelector":{"PlayerMode":"TARGET","PlayerCondition":{"IntCondition":{"IntProperty":"HANDSIZE","IntComparator":"GREATER","IntValue":1}},"Zone":"FIELD","CardMode":"TARGET","CardCondition":{"IntCondition":{"IntProperty":"HEALTH","IntComparator":"GREATER","IntValue":1}}},"EntityManipulations":[{"EntityIntManipulation":{"IntProperty":"HEALTH","IntOperator":"SUBTRACT","IntValue":1}}]}}}]}}],"Health":2,"Attack":1}}
' 'ne' --from $(cscli keys show alice --address) --chain-id testCardchain
