
import numpy as np
import matplotlib.pyplot as plt
x = np.array([pow(2, 4),	pow(2, 5),	pow(2, 6),	pow(
    2, 7),	pow(2, 8),	pow(2, 9), pow(2, 10)])
y = np.array([2.14248,	2.15303,	2.14166,	2.14512,	2.14184,	2.14160, 2.14142])
plt.title("issuing time diagram")
plt.plot(x, y, '-o')
plt.grid()

plt.show()


x = np.array([pow(2, 4),	pow(2, 5),	pow(2, 6),	pow(
    2, 7),	pow(2, 8),	pow(2, 9), pow(2, 10)])
y = np.array([4.14365,	5.961865,	9.659636,	17.102032,	32.155559,	61.493748,	120.608905])
plt.title("upload time diagram")
plt.plot(x, y, '-o')
plt.grid()
plt.show()
