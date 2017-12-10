#!/bin/bash

#========================
DATASETDIR=dataset
# input file is ${TYPE}0${INDEX}.txt ex) R101.txt
RESULTDIR=results
POPULATION=100
GENERATION=100
SELECTION=wsum       # pareto, wsum, ranksum
#CROSSOVER=bcrc      # uox, pmx, rc, bcrc
MUTATION=inversion
W_NVWHICLE=100
W_DISTANCE=0.001
ELITE=0
TOURNAMENT=3
CXRATE=0.6
MURATE=0.2
MUIRATE=0.03
SUFFIX=
#========================

if [ $# -ne 2 ]; then
    echo "[usage] $0 [uox, pmx, bcrc] cxrate" 1>&2
    exit 1
fi
CROSSOVER=${1}      # uox, pmx, bcrc
CXRATE=${2}

cd ../

for TYPE in R1 C1 RC1 R2 C2 RC2
do
    for INDEX in `seq 1 3`
    do
        for MURATE in 0.1 0.2 0.3 0.4
        do
            for i in `seq 1 10`
            do
                FILE=${TYPE}0${INDEX}
                INPUT=${DATASETDIR}/${TYPE}/${FILE}.txt
                OUTPUT=${RESULTDIR}/ex1/${CROSSOVER}/${FILE}/${CXRATE}

                mkdir -p ${OUTPUT}

                echo "${TYPE}0${INDEX}"
                ./main ${INPUT} ${OUTPUT} ${POPULATION} \
                    ${GENERATION} ${SELECTION} ${CROSSOVER} ${MUTATION} \
                    ${W_NVWHICLE} ${W_DISTANCE} ${ELITE} ${TOURNAMENT} \
                    ${CXRATE} ${MURATE} ${i}
            done
        done
    done
done
