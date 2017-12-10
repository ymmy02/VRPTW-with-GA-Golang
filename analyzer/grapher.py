import sys
import glob
import numpy as np
from scipy import stats
import matplotlib.pyplot as plt

def main(dirname):

    ALPHA = 0.95

    dirname = dirname + "/"
    ge = []
    nvavg = []
    nvbst = []
    diavg = []
    dibst = []

    ##############
    # Input Data #
    ##############
    files = glob.glob(dirname + "output*")
    for filename in files:
        finput = open(filename, 'r')
        line = finput.readline()
        while line:
            values = line.split(" ")
            generation = int(values[0].strip())
            nvehicle_avg = float(values[1].strip())
            distance_avg = float(values[2].strip())
            nvehicle_bst = float(values[3].strip())
            distance_bst = float(values[4].strip())
            ge.append(generation)
            nvavg.append(nvehicle_avg)
            diavg.append(distance_avg)
            nvbst.append(nvehicle_bst)
            dibst.append(distance_bst)
            line = finput.readline()
        finput.close()

    ge = np.array(ge).reshape(len(files), -1)
    nvavg = np.array(nvavg).reshape(len(files), -1)
    diavg = np.array(diavg).reshape(len(files), -1)
    nvbst = np.array(nvbst).reshape(len(files), -1)
    dibst = np.array(dibst).reshape(len(files), -1)

    nvavg_mean = np.mean(nvavg, axis=0)
    diavg_mean = np.mean(diavg, axis=0)
    nvbst_mean = np.mean(nvbst, axis=0)
    dibst_mean = np.mean(dibst, axis=0)

    nvavg_sem = stats.sem(nvavg)
    diavg_sem = stats.sem(diavg)
    nvbst_sem = stats.sem(nvbst)
    dibst_sem = stats.sem(dibst)

    nvavg_ci = stats.t.interval(ALPHA, len(files)-1, loc=nvavg_mean, scale=nvavg_sem)
    diavg_ci = stats.t.interval(ALPHA, len(files)-1, loc=diavg_mean, scale=diavg_sem)
    nvbst_ci = stats.t.interval(ALPHA, len(files)-1, loc=nvbst_mean, scale=nvbst_sem)
    dibst_ci = stats.t.interval(ALPHA, len(files)-1, loc=dibst_mean, scale=dibst_sem)

    # Avarage Num of Vehicles by Generations
    filename = "vehicles_avg_s.png"
    title = "Avarage Number of Vehicles by Generations"
    plot(dirname+filename, title, ge[0], nvavg_mean,  nvavg_ci, \
            [], [], "Genration", "Number of Vehicles", \
            "Vehicles", "")

    # Avarage Distance by Generations
    filename = "distance_avg_s.png"
    title = "Avarage Distance by Generations"
    plot(dirname+filename, title, ge[0], diavg_mean,  diavg_ci, \
            [], [], "Genration", "Distance", \
            "Distance", "")

    # Best Num of Vehicles by Generations
    filename = "vehicles_best_s.png"
    title = "Least Number of Vehicles by Generations"
    plot(dirname+filename, title, ge[0], nvbst_mean,  nvbst_ci, \
            [], [], "Genration", "Number of Vehicles", \
            "Vehicles", "")

    # Best Distance by Generations
    filename = "distance_best_s.png"
    title = "Least Distance by Generations"
    plot(dirname+filename, title, ge[0], dibst_mean,  dibst_ci, \
            [], [], "Genration", "Distance", \
            "Vehicles", "")

    # Num of Vehicles by Generations
    filename = "vehicles_s.png"
    title = "Number of Vehicles by Generations"
    plot(dirname+filename, title, ge[0], nvavg_mean,  nvavg_ci, \
            nvbst_mean, nvbst_ci, "Genration", "Number of Vehicles", \
            "Avarage", "Best")

    # Distance by Generations
    filename = "distance_s.png"
    title = "Distance by Generations"
    plot(dirname+filename, title, ge[0], diavg_mean,  diavg_ci, \
            dibst_mean, dibst_ci, "Genration", "Distance", \
            "Avarage", "Best")

def plot(filename, title, x, y1, y1e, y2, \
        y2e, xlabel, ylabel, label1, label2):
    plt.clf()
    plt.title(title)
    plt.xlabel(xlabel)
    plt.ylabel(ylabel)
    plt.plot(x, y1, label=label1, color='b')
    plt.plot(x, y1e[0], color='c')
    plt.plot(x, y1e[1], color='c')
    #plt.errorbar(x, y1, yerr=y1e, fmt='ro', ecolor='g')
    #plt.errorbar(x, y1, yerr=y1e)
    if len(y2) != 0:
        plt.plot(x, y2, label=label2, color='r')
        plt.plot(x, y2e[0], color='m')
        plt.plot(x, y2e[1], color='m')
        #plt.plot(x, y2, label=label2)
        #plt.errorbar(x, y2, yerr=y2e, fmt='ro', ecolor='g')
        #plt.errorbar(x, y2, yerr=y2e)
        plt.legend()
    plt.savefig(filename)

if __name__ == "__main__":

    args = sys.argv
    dirname = args[1]
    main(dirname)
