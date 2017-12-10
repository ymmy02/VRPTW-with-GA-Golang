import sys
import glob
from scipy import stats
from statistics import mean

def main(dirname):

    Alpha = 0.95    # 95% confidence
    dirname = dirname + "/"
    nvehicles = []
    distances = []

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
            nvehicles.append(nvehicle)
            distances.append(distance)
            line = finput.readline()
        finput.close()

    ##########################
    # Statistical Processing #
    ##########################
    nvmean = mean(nvehicles)
    nvsem = stats.sem(nvehicles)
    dimean = mean(distances)
    disem = stats.sem(distances)
    nvci = stats.t.interval(Alpha, len(nvehicles)-1, loc=nvmean, scale=nvsem)
    nverror = nvci[1] - nvmean
    dici = stats.t.interval(Alpha, len(distances)-1, loc=dimean, scale=disem)
    dierror = dici[1] - dimean

    ###############
    # Output Data #
    ###############
    foutput = open(dirname + "stat_analysis.txt", 'w')
    foutput.write("nvehicles 95%conf distances 95%conf\n")
    foutput.write(str(nvmean) + " " + str(nverror) + " " + str(dimean) + " " + str(dierror))
    foutput.close()

if __name__ == "__main__":

    args = sys.argv
    dirname = args[1]
    main(dirname)
