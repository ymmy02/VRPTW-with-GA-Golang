import sys
import glob
from scipy import stats
from statistics import mean

def main(dirname):

    Alpha = 0.95    # 95% confidence
    dirname = dirname + "/"
    nvehicles = []
    distances = []
    fitnesses = []

    ##############
    # Input Data #
    ##############
    files = glob.glob(dirname + "best_solutions*")
    for filename in files:
        finput = open(filename, 'r')
        line = finput.readline()
        while line:
            values = line.split(" ")
            nvehicle = int(values[0].strip())
            distance = float(values[1].strip())
            fitness = float(values[2].strip())
            nvehicles.append(nvehicle)
            distances.append(distance)
            fitnesses.append(fitnesses)
            line = finput.readline()
        finput.close()

    ##########################
    # Statistical Processing #
    ##########################
    nvmean = mean(nvehicles)
    nvsem = stats.sem(nvehicles)
    dimean = mean(distances)
    disem = stats.sem(distances)
    fimean = mean(fitnesses)
    fisem = stats.sem(fitnesses)
    nvci = stats.t.interval(Alpha, len(nvehicles)-1, loc=nvmean, scale=nvsem)
    nverror = nvci[1] - nvmean
    dici = stats.t.interval(Alpha, len(distances)-1, loc=dimean, scale=disem)
    dierror = dici[1] - dimean
    fici = stats.t.interval(Alpha, len(fitnesses)-1, loc=dimean, scale=disem)
    fierror = dici[1] - fimean

    ###############
    # Output Data #
    ###############
    foutput = open(dirname + "stat_analysis.txt", 'w')
    #foutput.write("nvehicles 95%conf distances 95%conf\n")
    #foutput.write(str(nvmean) + " " + str(nverror) + " " + str(dimean) + " " + str(dierror))
    foutput.write("{0:2.1f}".format(nvmean) + " " + "{0:2.2f}".format(nverror) + " " + "{0:4.0f}".format(dimean) + " " + "{0:3.0f}".format(dierror) + " " + "{0:3.1f}".format(fimean) + " " + "{0:3.1f}".format(fierror))
    foutput.close()

if __name__ == "__main__":

    args = sys.argv
    dirname = args[1]
    print(dirname)
    main(dirname)
