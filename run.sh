#!/bin/sh

#========================
DATASETDIR=dataset
TYPE=R1
INDEX=1
# input file is ${TYPE}0${INDEX}.txt ex) R101.txt
RESULTDIR=results
POPULATION=100
GENERATION=100
SELECTION=wsum    # pareto, wsum, ranksum
CROSSOVER=uox      # uox, pmx, rc, bcrc
MUTATION=inversion  # inversion, insersion
W_NVWHICLE=100
W_DISTANCE=0.001
ELITE=0
TOURNAMENT=3
CXRATE=0.6
MURATE=0.2
SUFFIX=
#========================

FILE=${TYPE}0${INDEX}
INPUT=${DATASETDIR}/${TYPE}/${FILE}.txt
OUTPUT=${RESULTDIR}/${TYPE}/${FILE}

mkdir -p ${OUTPUT}

go build main.go

./main ${INPUT} ${OUTPUT} ${POPULATION} \
    ${GENERATION} ${SELECTION} ${CROSSOVER} ${MUTATION} \
    ${W_NVWHICLE} ${W_DISTANCE} ${ELITE} ${TOURNAMENT} \
    ${CXRATE} ${MURATE} ${SUFFIX}
