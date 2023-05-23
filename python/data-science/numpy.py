import numpy as np

weights = np.array([0.5, 0.1, 1])
konto = np.array([1.54, 4.32, 10.5])
estimation = np.dot(weights, konto)
print(estimation)
