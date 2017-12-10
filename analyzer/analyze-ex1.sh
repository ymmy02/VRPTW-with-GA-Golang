#!/bin/sh

RESULTDIR=results

cd ../

for TYPE in C1 RC1 R2 C2 RC2
do
    for INDEX in `seq 1 3`
    do
        for CROSSOVER in uox pmx bcrc
        do
            for CXRATE in 0.6 0.7 0.8 0.9
            do
                DIR=${RESULTDIR}/ex1/wsum/${CROSSOVER}/${FILE}/${CXRATE}/${MURATE}
                python3 stat.py ${DIR}
                python3 grapher.py ${DIR}
            done
        done
    done
done
