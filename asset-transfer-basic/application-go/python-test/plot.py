
import numpy as np
import matplotlib.pyplot as plt
x = np.array([pow(2, 4),	pow(2, 5),	pow(2, 6),	pow(
    2, 7),	pow(2, 8),	pow(2, 9), pow(2, 10)])
y = np.array([2.139577,	2.138096,	2.138821,	2.138816,
             2.138568,	2.138432,	2.138484])
plt.title("issuing time diagram")
plt.plot(x, y, '-o')
plt.grid()

plt.show()


x = np.array([pow(2, 4),	pow(2, 5),	pow(2, 6),	pow(
    2, 7),	pow(2, 8),	pow(2, 9), pow(2, 10)])
y = np.array([4.027577,	5.845932,	9.538707,	16.862282,
             31.388436,	60.442806, 118.561657])
plt.title("upload time diagram")
plt.plot(x, y, '-o')
plt.grid()
plt.show()
