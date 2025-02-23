# Mandelbrot Set

The project generates image based on the mandelbrot set numbers.

The main idea for each point in 2D space we can iterate function based on coordinates until it gets out of predefined value of X.
It can also never pass X value, in that case we just defined a max iteration value.

## How to adjust?
The basic parameters are defined as constants, they are

- BOUNDARY, default is 2. Arbitary threshold.
- STEP, default is .0005. The step counter.

Above two values basically defines a matrix. For example, for the simplicity lets, take BOUNDARY as 2 and STEP as 0.5. Then the matrix will be.

$$
\begin{pmatrix}
-2.0, 2.0 & -1.5, 2.0 & -1.0, 2.0 & -0.5, 2.0 &  0.0, 2.0 &  0.5, 2.0 &  1.0, 2.0 &  1.5, 2.0 \\
-2.0, 1.5 & -1.5, 1.5 & -1.0, 1.5 & -0.5, 1.5 &  0.0, 1.5 &  0.5, 1.5 &  1.0, 1.5 &  1.5, 1.5 \\
-2.0, 1.0 & -1.5, 1.0 & -1.0, 1.0 & -0.5, 1.0 &  0.0, 1.0 &  0.5, 1.0 &  1.0, 1.0 &  1.5, 1.0 \\
-2.0, 0.5 & -1.5, 0.5 & -1.0, 0.5 & -0.5, 0.5 &  0.0, 0.5 &  0.5, 0.5 &  1.0, 0.5 &  1.5, 0.5 \\
-2.0, 0.0 & -1.5, 0.0 & -1.0, 0.0 & -0.5, 0.0 &  0.0, 0.0 &  0.5, 0.0 &  1.0, 0.0 &  1.5, 0.0 \\
-2.0, -0.5 & -1.5, -0.5 & -1.0, -0.5 & -0.5, -0.5 &  0.0, -0.5 &  0.5, -0.5 &  1.0, -0.5 &  1.5, -0.5 \\
-2.0, -1.0 & -1.5, -1.0 & -1.0, -1.0 & -0.5, -1.0 &  0.0, -1.0 &  0.5, -1.0 &  1.0, -1.0 &  1.5, -1.0 \\
-2.0, -1.5 & -1.5, -1.5 & -1.0, -1.5 & -0.5, -1.5 &  0.0, -1.5 &  0.5, -1.5 &  1.0, -1.5 &  1.5, -1.5 \\
\end{pmatrix}
$$

	//after this value we will break calculations because this means point never diverges from the set
- MANDELBROTMAXITERATION, default is 50. The max iteration number for never diverging elements.
- IMAGESCALEUP, default is 1. The scale up degree of each matrix point in photo, For large matrixes 1 is already ok.

## What it does?
Just generates cool images:)

## References
Basically below document's explanation converted to go code.
- https://www.wikihow.com/Plot-the-Mandelbrot-Set-By-Hand?utm_source=substack&utm_medium=email

