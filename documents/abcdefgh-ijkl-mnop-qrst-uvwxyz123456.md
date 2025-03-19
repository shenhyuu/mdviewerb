# 高等数学知识点 📚

### 一、函数、极限、连续 📉

#### 常见函数图像与性质 📊

- **幂函数**：$y = x^a$，定义域根据$a$不同而不同，具有加法、乘法等性质。
- **指数函数**：$y = a^x$（$a > 0$且$a \neq 1$），图像恒过点$(0,1)$，单调性取决于$a$的大小。
- **对数函数**：$y = \log_a x$（$a > 0$且$a \neq 1$），与指数函数互为反函数，图像恒过点$(1,0)$。
- **三角函数**：如$y = \sin x$、$y = \cos x$、$y = \tan x$等，具有周期性、奇偶性等性质。
- **反三角函数**：如$y = \arcsin x$、$y = \arccos x$、$y = \arctan x$等，定义域有限，具有特定的值域。

#### 极限 🔄

- **极限的定义**：设函数$f(x)$在点$x_0$的某个去心邻域内有定义，如果存在常数$A$，对于任意给定的正数$\varepsilon$，都存在正数$\delta$，使得当$0 < |x - x_0| < \delta$时，有$|f(x) - A| < \varepsilon$，则称$A$是$f(x)$当$x$趋近于$x_0$时的极限，记作$\lim_{x \to x_0} f(x) = A$。
- **极限的性质**：唯一性、局部有界性、局部保号性、不等式性等。
- **极限的存在准则**：
  - 夹逼定理：若在$x$的某个邻域内（$x = x_0$或$x \to \infty$）有$g(x) \leq f(x) \leq h(x)$，且$\lim_{x \to x_0} g(x) = \lim_{x \to x_0} h(x) = A$，则$\lim_{x \to x_0} f(x) = A$。
  - 单调有界定理：单调有界的数列必有极限。
- **两个重要极限**：
  - $\lim_{x \to 0} \frac{\sin x}{x} = 1$
  - $\lim_{x \to \infty} \left(1 + \frac{1}{x}\right)^x = e$或$\lim_{x \to 0} (1 + x)^{\frac{1}{x}} = e$

#### 连续 📈

- **连续的定义**：设函数$f(x)$在$x_0$的某个邻域内有定义，若$\lim_{x \to x_0} f(x) = f(x_0)$，则称$f(x)$在$x_0$处连续。
- **连续函数的性质**：四则运算性质、复合函数的连续性、反函数的连续性等。
- **闭区间上连续函数的性质**：
  - 有界性与最大值最小值定理：在闭区间上连续的函数在该区间上有界且能取得最大值和最小值。
  - 介值定理：设函数$f(x)$在区间$I$上连续，且$a$、$b$为$f(x)$在$I$上的两个函数值，则对于$a$和$b$之间的任意一个数$c$，存在$\xi \in I$，使得$f(\xi) = c$。
  - 零点定理：设函数$f(x)$在闭区间$[a, b]$上连续，且$f(a)$与$f(b)$异号，则在$(a, b)$内至少存在一个$\xi$，使得$f(\xi) = 0$。

### 二、一元函数微分学 📝

#### 导数与微分 ✏️

- **导数的定义**：设函数$y = f(x)$在点$x_0$的某个邻域内有定义，当自变量$x$在$x_0$处取得增量$\Delta x$（点$x_0 + \Delta x$仍在该邻域内）时，相应地函数取得增量$\Delta y = f(x_0 + \Delta x) - f(x_0)$。若极限$\lim_{\Delta x \to 0} \frac{\Delta y}{\Delta x}$存在，则称函数$y = f(x)$在点$x_0$处可导，称该极限为函数$y = f(x)$在点$x_0$处的导数，记作$f'(x_0)$或$y'|_{x=x_0}$。
- **导数的几何意义**：函数$y = f(x)$在点$x_0$处的导数$f'(x_0)$是曲线$y = f(x)$在点$(x_0, f(x_0))$处的切线斜率。
- **导数的基本公式**：
  - 常函数：$(C)' = 0$
  - 幂函数：$(x^a)' = a x^{a- 1}$
  - 指数函数：$(a^x)' = a^x \ln a$，$(e^x)' = e^x$
  - 对数函数：$(\log_a x)' = \frac{1}{x \ln a}$，$(\ln x)' = \frac{1}{x}$
  - 三角函数：$(\sin x)' = \cos x$，$(\cos x)' = - \sin x$，$(\tan x)' = \sec^2 x$，$(\cot x)' = - \csc^2 x$，$(\sec x)' = \sec x \tan x$，$(\csc x)' = - \csc x \cot x$
  - 反三角函数：$(\arcsin x)' = \frac{1}{\sqrt{1 - x^2}}$，$(\arccos x)' = - \frac{1}{\sqrt{1 - x^2}}$，$(\arctan x)' = \frac{1}{1 + x^2}$，$(\text{arccot}\, x)' = - \frac{1}{1 + x^2}$
- **导数的四则运算法则**：
  - $(u \pm v)' = u' \pm v'$
  - $(u \cdot v)' = u' \cdot v + u \cdot v'$
  - $\left( \frac{u}{v} \right)' = \frac{u' \cdot v - u \cdot v'}{v^2}$（$v \neq 0$）
- **复合函数的导数**：若$u = g(x)$在点$x$处可导，且$y = f(u)$在点$u$处可导，则复合函数$y = f(g(x))$在点$x$处可导，且其导数为$\frac{dy}{dx} = f'(u) \cdot g'(x)$。
- **微分的定义**：设函数$y = f(x)$在点$x_0$的某个邻域内有定义，若增量$\Delta y = f(x_0 + \Delta x) - f(x_0)$可表示为$\Delta y = A \Delta x + o(\Delta x)$，其中$A$是与$\Delta x$无关的常数，则称函数$y = f(x)$在点$x_0$处可微，称$A \Delta x$为函数$y = f(x)$在点$x_0$处的微分，记作$dy$或$df(x_0)$，即$dy = A \Delta x$。对于一元函数，可导与可微是等价的，且$A = f'(x_0)$，故微分也可表示为$dy = f'(x_0) \Delta x$。
- **微分的几何意义**：函数$y = f(x)$在点$x_0$处的微分$dy$是曲线$y = f(x)$在点$(x_0, f(x_0))$处的切线的纵坐标增量，即当$\Delta x$很小时，$\Delta y \approx dy$。
- **微分的基本公式**：与导数基本公式相对应，如$d(C) = 0$，$d(x^a) = a x^{a- 1} dx$，$d(e^x) = e^x dx$，$d(\sin x) = \cos x dx$等。
- **微分的四则运算法则**：与导数的四则运算法则类似，如$d(u \pm v) = du \pm dv$，$d(u \cdot v) = u dv + v du$，$d\left( \frac{u}{v} \right) = \frac{v du - u dv}{v^2}$（$v \neq 0$）。

#### 微分中值定理 ⚖️

- **罗尔定理**：若函数$f(x)$满足以下条件：
  - 在闭区间$[a, b]$上连续；
  - 在开区间$(a, b)$内可导；
  - $f(a) = f(b)$，

  则在$(a, b)$内至少存在一点$\xi$，使得$f'(\xi) = 0$。

- **拉格朗日中值定理**：若函数$f(x)$满足以下条件：
  - 在闭区间$[a, b]$上连续；
  - 在开区间$(a, b)$内可导，

  则在$(a, b)$内至少存在一点$\xi$，使得$f(b) - f(a) = f'(\xi)(b - a)$。

- **柯西中值定理**：若函数$f(x)$和$F(x)$满足以下条件：
  - 在闭区间$[a, b]$上连续；
  - 在开区间$(a, b)$内可导；
  - 对任一$x \in (a, b)$，有$F'(x) \neq 0$，

  则在$(a, b)$内至少存在一点$\xi$，使得$\frac{f(b) - f(a)}{F(b) - F(a)} = \frac{f'(\xi)}{F'(\xi)}$。

#### 导数的应用 🎯

- **洛必达法则**：在一定条件下，通过求导数的极限来解决$\frac{0}{0}$型、$\frac{\infty}{\infty}$型等不定型极限的问题。具体来说，若$\lim_{x \to x_0} f(x) = \lim_{x \to x_0} g(x) = 0$或$\lim_{x \to x_0} |f(x)| = \lim_{x \to x_0} |g(x)| = \infty$，且$f(x)$和$g(x)$在$x_0$的某去心邻域内可导，$g'(x) \neq 0$，且$\lim_{x \to x_0} \frac{f'(x)}{g'(x)}$存在（或为无穷大），则$\lim_{x \to x_0} \frac{f(x)}{g(x)} = \lim_{x \to x_0} \frac{f'(x)}{g'(x)}$。对于$x \to \infty$的情况类似。
- **泰勒公式**：设函数$f(x)$在点$x_0$处有$n$阶导数，则存在$x_0$的一个邻域，对于该邻域中的$x$，有
  $$
  f(x) = f(x_0) + f'(x_0)(x - x_0) + \frac{f''(x_0)}{2!}(x - x_0)^2 + \cdots + \frac{f^{(n)}(x_0)}{n!}(x - x_0)^n + R_n(x),
  $$
  其中，余项$R_n(x)$可以是佩亚诺余项$o((x - x_0)^n)$或拉格朗日余项$\frac{f^{(n+1)}(\xi)}{(n+1)!}(x - x_0)^{n+1}$（其中$\xi$在$x_0$和$x$之间）。当$x_0 = 0$时，泰勒公式称为麦克劳林公式。
- **函数单调性的判别**：设函数$f(x)$在区间$I$上可导，则$f(x)$在$I$上单调递增（或递减）的充要条件是$f'(x) \geq 0$（或$f'(x) \leq 0$）在$I$上恒成立。若在区间$I$内$f'(x) > 0$（或$f'(x) < 0$），则$f(x)$在$I$上严格单调递增（或递减）。
- **函数的极值**：
  - **定义**：设函数$f(x)$在点$x_0$的某个邻域内有定义，若对于该邻域内任意$x \neq x_0$，都有$f(x) < f(x_0)$（或$f(x) > f(x_0)$），则称$x_0$是函数$f(x)$的极大值点（或极小值点），对应的$f(x_0)$称为极大值（或极小值）。极大值和极小值统称为极值。
  - **费马定理**：若函数$f(x)$在点$x_0$处可导，且$x_0$是$f(x)$的极值点，则必有$f'(x_0) = 0$。满足$f'(x) = 0$的点称为函数的驻点，但驻点不一定是极值点。
  - **极值存在的充分条件**：
    - **第一充分条件**：设函数$f(x)$在$x_0$处连续，在$x_0$的某去心邻域内可导。若$f'(x)$在$x_0$处的左侧邻域内恒大于零，右侧邻域内恒小于零，则$x_0$是极大值点；若$f'(x)$在$x_0$处的左侧邻域内恒小于零，右侧邻域内恒大于零，则$x_0$是极小值点；若$f'(x)$在$x_0$的两侧邻域内符号不变，则$x_0$不是极值点。
    - **第二充分条件**：设函数$f(x)$在$x_0$处二阶可导，且$f'(x_0) = 0$，$f''(x_0) \neq 0$。若$f''(x_0) < 0$，则$x_0$是极大值点；若$f''(x_0) > 0$，则$x_0$是极小值点。
- **函数的最大值和最小值**：在闭区间$[a, b]$上连续的函数$f(x)$必有最大值和最小值。求最大值和最小值的方法是：求出函数在区间内的所有驻点和不可导点，计算这些点及区间端点处的函数值，比较这些函数值的大小，最大的即为最大值，最小的即为最小值。
- **曲线的凹凸性与拐点**：
  - **凹凸性的定义**：设函数$f(x)$在区间$I$上连续，若对于任意$x_1, x_2 \in I$及$\lambda \in (0, 1)$，有$f(\lambda x_1 + (1 - \lambda) x_2) < \lambda f(x_1) + (1 - \lambda) f(x_2)$（或$f(\lambda x_1 + (1 - \lambda) x_2) > \lambda f(x_1) + (1 - \lambda) f(x_2)$），则称$f(x)$在$I$上是凹的（或凸的）。
  - **凹凸性的判别**：若函数$f(x)$在区间$I$上二阶可导，则当$f''(x) > 0$时，$f(x)$在$I$上是凹的；当$f''(x) < 0$时，$f(x)$在$I$上是凸的。
  - **拐点的定义**：连续曲线$y = f(x)$上凹凸性发生变化的点称为拐点。
  - **拐点的求法**：求出函数$f(x)$的二阶导数$f''(x)$，解方程$f''(x) = 0$或找到$f''(x)$不存在的点，然后检验这些点两侧$f''(x)$的符号是否发生变化，若发生变化，则该点为拐点。
- **函数图像的描绘**：综合运用函数的单调性、极值、凹凸性、拐点、渐近线等性质，结合函数图像上的特殊点（如与坐标轴的交点、极值点、拐点等），绘制出函数的大致图像。
- **曲率**：
  - **曲率的定义**：曲率是描述曲线在某一点处弯曲程度的量。对于平面曲线$y = f(x)$，在点$(x, y)$处的曲率$\kappa$定义为
    $$
    \kappa = \frac{|y''|}{(1 + (y')^2)^{3/2}}.
    $$
  - **曲率半径**：曲率的倒数称为曲率半径，记作$R$，即$R = \frac{1}{\kappa}$。

### 三、一元函数积分学 📐

#### 不定积分 🔄

- **原函数与不定积分的定义**：若函数$F(x)$在区间$I$上可导，且$F'(x) = f(x)$对任意$x \in I$成立，则称$F(x)$是$f(x)$在区间$I$上的一个原函数。函数$f(x)$在区间$I$上的全体原函数称为$f(x)$的不定积分，记作$\int f(x) dx$，其中$\int$称为积分号，$f(x)$称为被积函数，$x$称为积分变量。
- **不定积分的性质**：
  - $\int [af(x) + bg(x)] dx = a \int f(x) dx + b \int g(x) dx$（$a$、$b$为常数）。
  - $\frac{d}{dx} \int f(x) dx = f(x)$，即不定积分是求导运算的逆运算。
- **不定积分的基本公式**：
  - $\int x^\alpha dx = \frac{x^{\alpha + 1}}{\alpha + 1} + C$（$\alpha \neq - 1$）
  - $\int \frac{1}{x} dx = \ln |x| + C$
  - $\int a^x dx = \frac{a^x}{\ln a} + C$（$a > 0$且$a \neq 1$）
  - $\int e^x dx = e^x + C$
  - $\int \cos x dx = \sin x + C$
  - $\int \sin x dx = - \cos x + C$
  - $\int \sec^2 x dx = \tan x + C$
  - $\int \csc^2 x dx = - \cot x + C$
  - $\int \sec x \tan x dx = \sec x + C$
  - $\int \csc x \cot x dx = - \csc x + C$
  - $\int \frac{1}{\sqrt{a^2 - x^2}} dx = \arcsin \frac{x}{a} + C$或$- \arccos \frac{x}{a} + C$（$a > 0$）
  - $\int \frac{1}{a^2 + x^2} dx = \frac{1}{a} \arctan \frac{x}{a} + C$（$a > 0$）
  - $\int \frac{1}{x^2 - a^2} dx = \frac{1}{2a} \ln \left| \frac{x - a}{x + a} \right| + C$（$a \neq 0$）
  - $\int \frac{1}{\sqrt{x^2 \pm a^2}} dx = \ln |x + \sqrt{x^2 \pm a^2}| + C$
  - $\int \tan x dx = - \ln |\cos x| + C$
  - $\int \cot x dx = \ln |\sin x| + C$
  - $\int \sec x dx = \ln |\sec x + \tan x| + C$
  - $\int \csc x dx = \ln |\tan \frac{x}{2}| + C$
- **换元积分法**：
  - **第一类换元法（凑微分法）**：若$\int f(u) du = F(u) + C$，且$u = \varphi(x)$可导，则$\int f(\varphi(x)) \varphi'(x) dx = F(\varphi(x)) + C$。具体操作是将被积函数中的某些部分看作一个整体进行变量替换，使积分更容易计算。
  - **第二类换元法**：对于形如$\int f(x) dx$的积分，通过变量替换$x = \varphi(t)$，将积分转化为关于$t$的积分$\int f(\varphi(t)) \varphi'(t) dt$，若新的积分容易计算，则可求得原积分的结果。常见的第二类换元法有三角替换（如处理$\sqrt{a^2 - x^2}$时令$x = a \sin \theta$）、根式替换（如处理含有根号的表达式时，令根号内的表达式为新变量）等。
- **分部积分法**：若$u(x)$和$v(x)$在区间$I$上可导，且$u'(x)$和$v'(x)$的不定积分存在，则有分部积分公式：
  $$
  \int u(x) v'(x) dx = u(x) v(x) - \int u'(x) v(x) dx.
  $$
  通常选择$v'(x) dx$为容易积分的部分，而$u(x)$为在求导后变得更简单的部分。

#### 定积分 📏

- **定积分的定义**：设函数$f(x)$在区间$[a, b]$上有定义。将区间$[a, b]$分成$n$个小区间，分点为$a = x_0 < x_1 < x_2 < \cdots < x_n = b$，小区间长度为$\Delta x_i = x_i - x_{i- 1}$（$i = 1, 2, \ldots, n$）。在每个小区间$[x_{i- 1}, x_i]$上任取一点$\xi_i$，作和式$S_n = \sum_{i=1}^n f(\xi_i) \Delta x_i$。若当$\lambda = \max_{1 \leq i \leq n} \Delta x_i \to 0$时，和式$S_n$的极限存在且与分法及$\xi_i$的取法无关，则称此极限为函数$f(x)$在区间$[a, b]$上的定积分，记作$\int_a^b f(x) dx$，即
  $$
  \int_a^b f(x) dx = \lim_{\lambda \to 0} \sum_{i=1}^n f(\xi_i) \Delta x_i.
  $$
- **定积分的存在条件**：若函数$f(x)$在区间$[a, b]$上连续，则$f(x)$在$[a, b]$上可积；若函数$f(x)$在区间$[a, b]$上有界且只有有限个间断点，则$f(x)$在$[a, b]$上可积。
- **定积分的性质**：
  - **线性性质**：$\int_a^b [af(x) + bg(x)] dx = a \int_a^b f(x) dx + b \int_a^b g(x) dx$（$a$、$b$为常数）。
  - **区间可加性**：$\int_a^b f(x) dx = \int_a^c f(x) dx + \int_c^b f(x) dx$（$a < c < b$）。
  - **不等式性质**：若在区间$[a, b]$上$f(x) \geq 0$，则$\int_a^b f(x) dx \geq 0$；若在区间$[a, b]$上$f(x) \leq g(x)$，则$\int_a^b f(x) dx \leq \int_a^b g(x) dx$。
  - **绝对可积性**：若$f(x)$在区间$[a, b]$上可积，则$|f(x)|$在$[a, b]$上也可积，且$\left| \int_a^b f(x) dx \right| \leq \int_a^b |f(x)| dx$。
- **微积分基本定理**：
  - **变上限积分函数的导数**：若函数$f(x)$在区间$[a, b]$上连续，则函数$F(x) = \int_a^x f(t) dt$在$[a, b]$上可导，且$F'(x) = f(x)$。
  - **牛顿- 莱布尼茨公式**：若函数$f(x)$在区间$[a, b]$上连续，且$F(x)$是$f(x)$的一个原函数，则
    $$
    \int_a^b f(x) dx = F(b) - F(a).
    $$
- **定积分的换元法和分部积分法**：
  - **换元法**：若$f(x)$在区间$[a, b]$上连续，且$x = \varphi(t)$满足$\varphi(\alpha) = a$、$\varphi(\beta) = b$，且$\varphi(t)$在$[\alpha, \beta]$上单调、可导，导数不为零，则有
    $$
    \int_a^b f(x) dx = \int_\alpha^\beta f(\varphi(t)) \varphi'(t) dt.
    $$
  - **分部积分法**：对于定积分$\int_a^b f(x) g'(x) dx$，若$f(x)$和$g(x)$在$[a, b]$上可导，则有
    $$
    \int_a^b f(x) g'(x) dx = [f(x) g(x)]_a^b - \int_a^b f'(x) g(x) dx.
    $$
- **反常积分（广义积分）**：
  - **无穷限反常积分**：对于积分$\int_a^\infty f(x) dx$，定义为$\lim_{b \to \infty} \int_a^b f(x) dx$，若该极限存在，则称反常积分收敛，否则发散。类似地，可定义$\int_{- \infty}^b f(x) dx$和$\int_{- \infty}^\infty f(x) dx$。
  - **无界函数的反常积分**：对于积分$\int_a^b f(x) dx$，若$f(x)$在$x = c$（$a < c < b$）处无界，则定义为$\lim_{t \to c^- } \int_a^t f(x) dx + \lim_{t \to c^+} \int_t^b f(x) dx$，若两个极限都存在，则反常积分收敛，否则发散。
  - **反常积分的收敛判别法**：
    - **比较判别法**：设在积分区间上$0 \leq f(x) \leq g(x)$，若$\int g(x) dx$收敛，则$\int f(x) dx$收敛；若$\int f(x) dx$发散，则$\int g(x) dx$发散。
    - **极限比较判别法**：设$\lim_{x \to \infty} \frac{f(x)}{g(x)} = L$，若$0 < L < \infty$，则$\int_a^\infty f(x) dx$和$\int_a^\infty g(x) dx$同敛散。
    - **p- 判别法**：对于无穷限反常积分$\int_a^\infty \frac{1}{x^p} dx$，当$p > 1$时收敛，$p \leq 1$时发散；对于无界函数的反常积分$\int_a^b \frac{1}{(x - c)^p} dx$，当$p < 1$时收敛，$p \geq 1$时发散。

#### 定积分的应用 🎯

- **几何应用**：
  - **平面图形的面积**：
    - 对于由曲线$y = f(x)$、$y = g(x)$（$f(x) \geq g(x)$）及直线$x = a$、$x = b$所围成的平面图形，其面积$A$为
      $$
      A = \int_a^b [f(x) - g(x)] dx.
      $$
    - 对于由参数方程$x = \varphi(t)$、$y = \psi(t)$（$\alpha \leq t \leq \beta$）所表示的曲线围成的图形，其面积$A$为
      $$
      A = \int_\alpha^\beta y \frac{dx}{dt} dt \quad \text{或} \quad A = \int_\alpha^\beta x \frac{dy}{dt} dt,
      $$
      具体取决于曲线的方向和形状。
    - 对于极坐标下的曲线$r = r(\theta)$（$\alpha \leq \theta \leq \beta$）所围成的图形，其面积$A$为
      $$
      A = \frac{1}{2} \int_\alpha^\beta r^2 d\theta.
      $$
  - **体积**：
    - **旋转体的体积**：将平面图形绕 x 轴或 y 轴旋转一周所得到的立体的体积。对于绕 x 轴旋转的旋转体，若平面图形由曲线$y = f(x)$、直线$x = a$、$x = b$及 x 轴围成，则体积$V$为
      $$
      V = \pi \int_a^b [f(x)]^2 dx.
      $$
      对于绕 y 轴旋转的旋转体，若平面图形由曲线$x = g(y)$、直线$y = c$、$y = d$及 y 轴围成，则体积$V$为
      $$
      V = \pi \int_c^d [g(y)]^2 dy.
      $$
    - **平行截面面积为已知的立体的体积**：若立体在高度$x$处的截面面积为$A(x)$，则体积$V$为
      $$
      V = \int_a^b A(x) dx.
      $$
  - **曲线的弧长**：
    - 对于直角坐标下的曲线$y = f(x)$（$a \leq x \leq b$），其弧长$L$为
      $$
      L = \int_a^b \sqrt{1 + [f'(x)]^2} dx.
      $$
    - 对于参数方程$x = \varphi(t)$、$y = \psi(t)$（$\alpha \leq t \leq \beta$）所表示的曲线，其弧长$L$为
      $$
      L = \int_\alpha^\beta \sqrt{\left( \frac{dx}{dt} \right)^2 + \left( \frac{dy}{dt} \right)^2} dt.
      $$
    - 对于极坐标下的曲线$r = r(\theta)$（$\alpha \leq \theta \leq \beta$），其弧长$L$为
      $$
      L = \int_\alpha^\beta \sqrt{\left( \frac{dr}{d\theta} \right)^2 + r^2} d\theta.
      $$
- **物理应用**：
  - **变力沿直线所做的功**：若物体从点$x = a$移动到点$x = b$的过程中，受到的力$F(x)$随位置$x$变化，则所做的功$W$为
    $$
    W = \int_a^b F(x) dx.
    $$
  - **液体的侧压力**：深度为$h$处的侧压力为$P = \rho g h$，其中$\rho$是液体的密度，$g$是重力加速度。对于一矩形薄板垂直沉入水中，其上沿在深度$a$处，下沿在深度$b$处，宽度为$w$，则薄板所受的侧压力$F$为
    $$
    F = \int_a^b \rho g h \cdot w \, dh.
    $$
  - **质心坐标**：对于平面薄片，其密度为$\rho(x, y)$，所占区域为$D$，则质心坐标$(\bar{x}, \bar{y})$为
    $$
    \bar{x} = \frac{1}{M} \iint_D x \rho(x, y) d\sigma, \quad \bar{y} = \frac{1}{M} \iint_D y \rho(x, y) d\sigma,
    $$
    其中$M = \iint_D \rho(x, y) d\sigma$为薄片的总质量。对于由曲线$y = f(x)$、直线$x = a$、$x = b$及 x 轴围成的均匀薄片（密度$\rho = 1$），其质心坐标为
    $$
    \bar{x} = \frac{1}{A} \int_a^b x [f(x) - 0] dx, \quad \bar{y} = \frac{1}{A} \int_a^b \frac{1}{2} [f(x)]^2 dx,
    $$
    其中$A = \int_a^b f(x) dx$为区域面积。

### 四、多元函数微积分学 📦

#### 多元函数的基本概念 📂

- **多元函数的定义**：设$D$是平面上的一个点集，若对于每个点$(x, y) \in D$，变量$z$按照一定的法则都有唯一确定的值与之对应，则称$z$是$x$和$y$的二元函数，记作$z = f(x, y)$，其中$x$、$y$称为自变量，$z$称为因变量，$D$称为定义域。类似地，可定义三元及更多元的函数。
- **二元函数的极限**：设函数$z = f(x, y)$在点$P_0(x_0, y_0)$的某个邻域内有定义（$P_0$可以是定义域内的点或聚点）。若存在常数$A$，对于任意给定的正数$\varepsilon$，都存在正数$\delta$，使得当$0 < |P - P_0| < \delta$（即$0 < \sqrt{(x - x_0)^2 + (y - y_0)^2} < \delta$）时，有$|f(x, y) - A| < \varepsilon$，则称$A$是函数$f(x, y)$当$(x, y) \to (x_0, y_0)$时的极限，记作$\lim_{(x, y) \to (x_0, y_0)} f(x, y) = A$。对于$(x, y) \to \infty$的情况，类似地定义极限。
- **二元函数的连续性**：设函数$z = f(x, y)$在点$P_0(x_0, y_0)$的某个邻域内有定义，若$\lim_{(x, y) \to (x_0, y_0)} f(x, y) = f(x_0, y_0)$，则称$f(x, y)$在点$P_0$处连续。若函数$f(x, y)$在定义域$D$上的每一点都不间断，则称$f(x, y)$在$D$上连续。
- **连续函数的性质**：多元连续函数在有界闭区域上具有有界性、最大值最小值定理、介值定理等性质。

#### 偏导数 📉

- **偏导数的定义**：设函数$z = f(x, y)$在点$(x_0, y_0)$的某个邻域内有定义，若极限
  $$
  \frac{\partial f}{\partial x}\bigg|_{(x_0, y_0)} = \lim_{\Delta x \to 0} \frac{f(x_0 + \Delta x, y_0) - f(x_0, y_0)}{\Delta x}
  $$
  存在，则称此极限为函数$z = f(x, y)$在点$(x_0, y_0)$处对$x$的偏导数，记作$f_x(x_0, y_0)$或$\frac{\partial z}{\partial x}\bigg|_{(x_0, y_0)}$。类似地，可定义对$y$的偏导数$f_y(x_0, y_0)$。
- **高阶偏导数**：若函数$z = f(x, y)$在区域$D$内对$x$的偏导数$f_x(x, y)$仍然可对$x$或$y$求偏导数，则得到二阶偏导数，记作$f_{xx}$、$f_{xy}$、$f_{yx}$、$f_{yy}$等。对于多元函数的高阶偏导数，一般情况下混合偏导数$f_{xy}$和$f_{yx}$在连续的条件下相等，即若$f_{xy}$和$f_{yx}$在点$(x_0, y_0)$的某个邻域内连续，则$f_{xy}(x_0, y_0) = f_{yx}(x_0, y_0)$。
- **偏导数的几何意义**：函数$z = f(x, y)$在点$(x_0, y_0)$处对$x$的偏导数$f_x(x_0, y_0)$是曲面$z = f(x, y)$在点$(x_0, y_0, f(x_0, y_0))$处的切平面沿着$x$轴方向的斜率，类似地，对$y$的偏导数是沿$y$轴方向的斜率。
- **方向导数**：设函数$u = f(x, y, z)$在点$P(x_0, y_0, z_0)$的某个邻域内有定义，$\mathbf{l}$是从点$P$指出的任一方向，其方向余弦为$\cos \alpha$、$\cos \beta$、$\cos \gamma$。若极限
  $$
  \frac{\partial f}{\partial \mathbf{l}}\bigg|_{P} = \lim_{h \to 0} \frac{f(P + h\mathbf{l}) - f(P)}{h}
  $$
  存在，则称此极限为函数$f$在点$P$处沿方向$\mathbf{l}$的方向导数，记作$\frac{\partial f}{\partial \mathbf{l}} \bigg|_{P}$。方向导数与梯度的关系为$\frac{\partial f}{\partial \mathbf{l}} = \nabla f \cdot \mathbf{l}$，其中$\nabla f$是函数$f$在该点的梯度向量。
- **梯度**：设函数$u = f(x, y, z)$在点$P(x_0, y_0, z_0)$处存在连续偏导数，则梯度向量$\nabla f$定义为
  $$
  \nabla f = \left( \frac{\partial f}{\partial x}, \frac{\partial f}{\partial y}, \frac{\partial f}{\partial z} \right).
  $$
  梯度的方向是函数在该点处增长最快的方向，其模长等于该方向的方向导数。

#### 全微分 📊

- **全微分的定义**：设函数$z = f(x, y)$在点$(x_0, y_0)$处的全增量为$\Delta z = f(x_0 + \Delta x, y_0 + \Delta y) - f(x_0, y_0)$。若存在常数$A$和$B$，使得
  $$
  \Delta z = A \Delta x + B \Delta y + o(\sqrt{(\Delta x)^2 + (\Delta y)^2}),
  $$
  则称函数$z = f(x, y)$在点$(x_0, y_0)$处可微，称线性主部$A \Delta x + B \Delta y$为函数$z = f(x, y)$在该点的全微分，记作$dz$或$df(x_0, y_0)$，即
  $$
  dz = A \Delta x + B \Delta y.
  $$
  对于二元函数，若其在点$(x_0, y_0)$处的偏导数$f_x$和$f_y$存在且连续，则函数在该点可微，且全微分为
  $$
  dz = f_x(x_0, y_0) \Delta x + f_y(x_0, y_0) \Delta y.
  $$
- **全微分的几何意义**：全微分$dz$是曲面$z = f(x, y)$在点$(x_0, y_0, z_0)$处的切平面的纵坐标增量，当$\Delta x$和$\Delta y$很小时，可以用全微分近似代替全增量，即$\Delta z \approx dz$。

#### 多元复合函数与隐函数的求导法则 🔄

- **链式法则**：
  - **一阶全导数**：若$z = f(u, v)$，而$u = \varphi(t)$、$v = \psi(t)$，则
    $$
    \frac{dz}{dt} = \frac{\partial f}{\partial u} \cdot \frac{du}{dt} + \frac{\partial f}{\partial v} \cdot \frac{dv}{dt}.
    $$
  - **一阶偏导数**：若$z = f(u, v)$，而$u = \varphi(x, y)$、$v = \psi(x, y)$，则
    $$
    \frac{\partial z}{\partial x} = \frac{\partial f}{\partial u} \cdot \frac{\partial u}{\partial x} + \frac{\partial f}{\partial v} \cdot \frac{\partial v}{\partial x},
    $$
    $$
    \frac{\partial z}{\partial y} = \frac{\partial f}{\partial u} \cdot \frac{\partial u}{\partial y} + \frac{\partial f}{\partial v} \cdot \frac{\partial v}{\partial y}.
    $$
- **隐函数的求导公式**：
  - **由方程$F(x, y) = 0$确定的隐函数$y = y(x)$**：
    $$
    \frac{dy}{dx} = - \frac{F_x}{F_y},
    $$
    其中$F_x$和$F_y$分别是$F$对$x$和$y$的偏导数，且$F_y \neq 0$。
  - **由方程组$F(x, y, u, v) = 0$、$G(x, y, u, v) = 0$确定的隐函数$u = u(x, y)$、$v = v(x, y)$**：
    $$
    \frac{\partial u}{\partial x} = - \frac{\begin{vmatrix} F_x & F_v \\ G_x & G_v \end{vmatrix}}{\begin{vmatrix} F_u & F_v \\ G_u & G_v \end{vmatrix}}, \quad
    \frac{\partial u}{\partial y} = - \frac{\begin{vmatrix} F_y & F_v \\ G_y & G_v \end{vmatrix}}{\begin{vmatrix} F_u & F_v \\ G_u & G_v \end{vmatrix}},
    $$
    $$
    \frac{\partial v}{\partial x} = - \frac{\begin{vmatrix} F_u & F_x \\ G_u & G_x \end{vmatrix}}{\begin{vmatrix} F_u & F_v \\ G_u & G_v \end{vmatrix}}, \quad
    \frac{\partial v}{\partial y} = - \frac{\begin{vmatrix} F_u & F_y \\ G_u & G_y \end{vmatrix}}{\begin{vmatrix} F_u & F_v \\ G_u & G_v \end{vmatrix}}.
    $$
    其中，分子中的行列式是由对应的偏导数组成的二阶行列式，分母行列式为雅可比行列式$J = \begin{vmatrix} F_u & F_v \\ G_u & G_v \end{vmatrix}$，且$J \neq 0$。

#### 多元函数的极值 🎯

- **多元函数极值的定义**：设函数$f(x, y)$在点$(x_0, y_0)$的某个邻域内有定义，若对于该邻域内任意$(x, y) \neq (x_0, y_0)$，都有$f(x, y) < f(x_0, y_0)$（或$f(x, y) > f(x_0, y_0)$），则称$(x_0, y_0)$是函数$f(x, y)$的极大值点（或极小值点），对应的$f(x_0, y_0)$称为极大值（或极小值）。极大值和极小值统称为极值。
- **多元函数极值存在的条件**：
  - **必要条件**：若函数$f(x, y)$在点$(x_0, y_0)$处有极值，且偏导数存在，则必有$f_x(x_0, y_0) = 0$、$f_y(x_0, y_0) = 0$。满足这两个条件的点称为函数的临界点或驻点，但驻点不一定是极值点。
  - **充分条件**：设函数$f(x, y)$在点$(x_0, y_0)$处有二阶连续偏导数，且$f_x(x_0, y_0) = 0$、$f_y(x_0, y_0) = 0$，令
    $$
    A = f_{xx}(x_0, y_0), \quad B = f_{xy}(x_0, y_0), \quad C = f_{yy}(x_0, y_0),
    $$
    则：
    - 当$AC - B^2 > 0$且$A < 0$（或$C < 0$）时，函数在该点处有极大值；
    - 当$AC - B^2 > 0$且$A > 0$（或$C > 0$）时，函数在该点处有极小值；
    - 当$AC - B^2 < 0$时，函数在该点处无极值；
    - 当$AC - B^2 = 0$时，无法确定，需进一步分析。
- **条件极值与拉格朗日乘数法**：在约束条件$\varphi(x, y) = 0$下，求函数$f(x, y)$的极值，可构造拉格朗日函数$L(x, y, \lambda) = f(x, y) + \lambda \varphi(x, y)$，解方程组
  $$
  \begin{cases}
  \frac{\partial L}{\partial x} = 0, \\
  \frac{\partial L}{\partial y} = 0, \\
  \frac{\partial L}{\partial \lambda} = 0,
  \end{cases}
  $$
  即可找到可能的极值点。对于多个约束条件的情况，类似地引入多个拉格朗日乘数。

#### 二重积分 📐

- **二重积分的定义**：设$f(x, y)$是定义在有界闭区域$D$上的函数。将区域$D$划分为$n$个子区域$\Delta D_1, \Delta D_2, \ldots, \Delta D_n$，每个子区域的面积记为$\Delta \sigma_i$。在每个子区域$\Delta D_i$上任取一点$(\xi_i, \eta_i)$，作和式$S_n = \sum_{i=1}^n f(\xi_i, \eta_i) \Delta \sigma_i$。若当各子区域的直径的最大值$\lambda \to 0$时，和式$S_n$的极限存在且与划分方式及点$(\xi_i, \eta_i)$的取法无关，则称此极限为函数$f(x, y)$在区域$D$上的二重积分，记作$\iint_D f(x, y) d\sigma$，即
  $$
  \iint_D f(x, y) d\sigma = \lim_{\lambda \to 0} \sum_{i=1}^n f(\xi_i, \eta_i) \Delta \sigma_i.
  $$
- **二重积分的性质**：
  - **线性性质**：$\iint_D [af(x, y) + bg(x, y)] d\sigma = a \iint_D f(x, y) d\sigma + b \iint_D g(x, y) d\sigma$（$a$、$b$为常数）。
  - **区域可加性**：若区域$D$分为两个子区域$D_1$和$D_2$，则$\iint_D f(x, y) d\sigma = \iint_{D_1} f(x, y) d\sigma + \iint_{D_2} f(x, y) d\sigma$。
  - **不等式性质**：若在区域$D$上$f(x, y) \geq 0$，则$\iint_D f(x, y) d\sigma \geq 0$；若在区域$D$上$f(x, y) \leq g(x, y)$，则$\iint_D f(x, y) d\sigma \leq \iint_D g(x, y) d\sigma$。
- **二重积分的计算**：
  - **直角坐标系下的计算**：对于$x$- 型区域（即由$x = a$、$x = b$、$y = \varphi_1(x)$、$y = \varphi_2(x)$围成的区域），二重积分可转化为累次积分：
    $$
    \iint_D f(x, y) d\sigma = \int_a^b \left( \int_{\varphi_1(x)}^{\varphi_2(x)} f(x, y) dy \right) dx.
    $$
    对于$y$- 型区域（即由$y = c$、$y = d$、$x = \psi_1(y)$、$x = \psi_2(y)$围成的区域），类似地有：
    $$
    \iint_D f(x, y) d\sigma = \int_c^d \left( \int_{\psi_1(y)}^{\psi_2(y)} f(x, y) dx \right) dy.
    $$
  - **极坐标系下的计算**：当积分区域$D$或被积函数$f(x, y)$具有圆的对称性或易于用极坐标表示时，可将二重积分转化为极坐标下的累次积分。设$x = r \cos \theta$、$y = r \sin \theta$，则面积元素$d\sigma = r dr d\theta$，积分区域需用极坐标表示。例如，对于区域$D$由$r = a(\theta)$围成，则
    $$
    \iint_D f(x, y) d\sigma = \int_{\theta_1}^{\theta_2} \int_0^{a(\theta)} f(r \cos \theta, r \sin \theta) r dr d\theta.
    $$
- **二重积分的应用**：
  - **平面薄片的质心**：对于密度为$\rho(x, y)$的平面薄片，其质心坐标为
    $$
    \bar{x} = \frac{1}{M} \iint_D x \rho(x, y) d\sigma, \quad \bar{y} = \frac{1}{M} \iint_D y \rho(x, y) d\sigma,
    $$
    其中$M = \iint_D \rho(x, y) d\sigma$为薄片的总质量。
  - **平面薄片的转动惯量**：对于密度为$\rho(x, y)$的平面薄片，其对 x 轴和 y 轴的转动惯量分别为
    $$
    I_x = \iint_D y^2 \rho(x, y) d\sigma, \quad I_y = \iint_D x^2 \rho(x, y) d\sigma.
    $$
  - **曲面的面积**：对于由$z = f(x, y)$给出的曲面，其在区域$D$上的面积$A$为
    $$
    A = \iint_D \sqrt{1 + \left( \frac{\partial f}{\partial x} \right)^2 + \left( \frac{\partial f}{\partial y} \right)^2} d\sigma.
    $$

### 五、无穷级数 🔄

#### 常数项级数 📊

- **常数项级数的定义**：将数列$\{a_n\}$的各项相加，得到无穷级数$\sum_{n=1}^\infty a_n = a_1 + a_2 + \cdots + a_n + \cdots$。若部分和数列$\{S_n\}$（其中$S_n = a_1 + a_2 + \cdots + a_n$）收敛于$S$，则称级数收敛，称$S$为级数的和；否则称级数发散。
- **收敛级数的性质**：
  - **线性性质**：若级数$\sum a_n$和$\sum b_n$收敛，且和分别为$S$和$T$，则级数$\sum (ka_n + lb_n)$（$k$、$l$为常数）也收敛，其和为$kS + lT$。
  - **添加或去掉有限项不影响级数的敛散性**。
  - **收敛级数的通项趋于零**：若级数$\sum a_n$收敛，则$\lim_{n \to \infty} a_n = 0$。
- **收敛性的判别法**：
  - **正项级数的判别法**：
    - **比较判别法**：设$\sum a_n$和$\sum b_n$都是正项级数，若存在自然数$N$，当$n \geq N$时有$a_n \leq b_n$，则$\sum b_n$收敛$\Rightarrow \sum a_n$收敛；$\sum a_n$发散$\Rightarrow \sum b_n$发散。
    - **极限比较判别法**：设$\sum a_n$和$\sum b_n$都是正项级数，且$\lim_{n \to \infty} \frac{a_n}{b_n} = L$（$0 < L < \infty$），则$\sum a_n$和$\sum b_n$同敛散。
    - **比值判别法（达朗贝尔判别法）**：对于正项级数$\sum a_n$，若$\lim_{n \to \infty} \frac{a_{n+1}}{a_n} = q$，则当$q < 1$时级数收敛，当$q > 1$时级数发散，当$q = 1$时判别法失效。
    - **根值判别法（柯西判别法）**：对于正项级数$\sum a_n$，若$\lim_{n \to \infty} \sqrt[n]{a_n} = q$，则当$q < 1$时级数收敛，当$q > 1$时级数发散，当$q = 1$时判别法失效。
    - **积分判别法**：设$f(x)$是定义在区间$[N, +\infty)$上的非负减函数，且当$x \geq N$时$f(x) \geq 0$且单调递减，则级数$\sum_{n=N}^\infty f(n)$与广义积分$\int_N^{+\infty} f(x) dx$同敛散。
  - **交错级数的判别法（莱布尼茨判别法）**：对于交错级数$\sum (- 1)^{n- 1} a_n$（$a_n > 0$），若$\{a_n\}$单调递减且$\lim_{n \to \infty} a_n = 0$，则级数收敛。
  - **绝对收敛与条件收敛**：若级数$\sum |a_n|$收敛，则称级数$\sum a_n$绝对收敛；若级数$\sum a_n$收敛，但$\sum |a_n|$发散，则称级数$\sum a_n$条件收敛。绝对收敛的级数必定收敛。

#### 幂级数 📊

- **幂级数的定义**：形如$\sum_{n=0}^\infty a_n (x - x_0)^n$的级数称为幂级数，其中$x$是变量，$x_0$是收敛中心，$a_n$是系数。特别地，当$x_0 = 0$时，幂级数为$\sum_{n=0}^\infty a_n x^n$，称为标准幂级数。
- **收敛半径与收敛区间**：对于幂级数$\sum_{n=0}^\infty a_n (x - x_0)^n$，存在一个非负实数$R$，使得当$|x - x_0| < R$时级数绝对收敛，当$|x - x_0| > R$时级数发散，称$R$为收敛半径。收敛区间的中心为$x_0$，半径为$R$，即$(x_0 - R, x_0 + R)$。在收敛区间的端点处，级数可能收敛也可能发散，需单独判断。
- **收敛半径的求法**：
  - **根值法**：$R = \frac{1}{\limsup_{n \to \infty} \sqrt[n]{|a_n|}}$。
  - **比值法**：若$\lim_{n \to \infty} \left| \frac{a_{n+1}}{a_n} \right| = q$，则$R = \frac{1}{q}$。若极限不存在，则根值法更有效。
- **幂级数的性质**：
  - **和函数的连续性**：幂级数在其收敛区间内和函数连续。
  - **逐项求导与逐项积分**：幂级数在其收敛区间内可逐项求导和逐项积分，即
    $$
    \frac{d}{dx} \sum_{n=0}^\infty a_n (x - x_0)^n = \sum_{n=1}^\infty n a_n (x - x_0)^{n- 1},
    $$
    $$
    \int \sum_{n=0}^\infty a_n (x - x_0)^n dx = C + \sum_{n=0}^\infty \frac{a_n}{n + 1} (x - x_0)^{n + 1},
    $$
    其中逐项求导后的收敛半径与原级数相同，逐项积分后的收敛半径也与原级数相同。
- **泰勒级数与麦克劳林级数**：若函数$f(x)$在点$x_0$的某个邻域内具有各阶导数，则可展开为泰勒级数
  $$
  f(x) = \sum_{n=0}^\infty \frac{f^{(n)}(x_0)}{n!} (x - x_0)^n,
  $$
  其中收敛半径由该点处的幂级数展开式确定。当$x_0 = 0$时，泰勒级数称为麦克劳林级数。常见的函数展开式如：
  - $e^x = \sum_{n=0}^\infty \frac{x^n}{n!}$（$|x| < \infty$）
  - $\sin x = \sum_{n=0}^\infty (- 1)^n \frac{x^{2n+1}}{(2n+1)!}$（$|x| < \infty$）
  - $\cos x = \sum_{n=0}^\infty (- 1)^n \frac{x^{2n}}{(2n)!}$（$|x| < \infty$）
  - $\ln(1 + x) = \sum_{n=1}^\infty (- 1)^{n- 1} \frac{x^n}{n}$（$|x| < 1$）
  - $\frac{1}{1 - x} = \sum_{n=0}^\infty x^n$（$|x| < 1$）
  - $\arctan x = \sum_{n=0}^\infty (- 1)^n \frac{x^{2n+1}}{2n+1}$（$|x| \leq 1$）

#### 傅里叶级数 📊

- **三角函数系的正交性**：在区间$[- l, l]$上，三角函数系$\{ \frac{1}{\sqrt{2l}}, \cos \frac{n\pi x}{l}, \sin \frac{n\pi x}{l} \mid n = 1, 2, \ldots \}$是正交函数系，即满足
  $$
  \int_{- l}^l \cos \frac{m\pi x}{l} \cos \frac{n\pi x}{l} dx = \begin{cases} 0, & m \neq n, \\ l, & m = n \neq 0, \\ 2l, & m = n = 0, \end{cases}
  $$
  $$
  \int_{- l}^l \sin \frac{m\pi x}{l} \sin \frac{n\pi x}{l} dx = \begin{cases} 0, & m \neq n, \\ l, & m = n, \end{cases}
  $$
  $$
  \int_{- l}^l \cos \frac{m\pi x}{l} \sin \frac{n\pi x}{l} dx = 0.
  $$
- **傅里叶系数与傅里叶级数**：对于定义在区间$[- l, l]$上的函数$f(x)$，其傅里叶级数为
  $$
  \frac{a_0}{2} + \sum_{n=1}^\infty \left( a_n \cos \frac{n\pi x}{l} + b_n \sin \frac{n\pi x}{l} \right),
  $$
  其中系数$a_n$和$b_n$称为傅里叶系数，计算公式为
  $$
  a_n = \frac{1}{l} \int_{- l}^l f(x) \cos \frac{n\pi x}{l} dx \quad (n = 0, 1, 2, \ldots),
  $$
  $$
  b_n = \frac{1}{l} \int_{- l}^l f(x) \sin \frac{n\pi x}{l} dx \quad (n = 1, 2, \ldots).
  $$
- **收敛定理（狄利克雷条件）**：若函数$f(x)$满足：
  - 在区间$[- l, l]$上$f(x)$按段光滑（即在区间内只有有限个极值点，且曲线只有有限个间断点，且在这些间断点处左、右极限都存在），
  - 在区间$[- l, l]$上$f(x)$有有限个极值点，

  则$f(x)$的傅里叶级数在区间$[- l, l]$上收敛。对于任意$x \in [- l, l]$，当$x$是$f(x)$的连续点时，傅里叶级数收敛于$f(x)$；当$x$是$f(x)$的间断点时，傅里叶级数收敛于$\frac{1}{2}[f(x^- ) + f(x^+)]$，其中$f(x^- )$和$f(x^+)$分别为$x$处的左极限和右极限。

### 六、常微分方程 📝

#### 基本概念 📂

- **微分方程的定义**：含有未知函数及其导数的方程称为微分方程。若未知函数是一元函数，则称为常微分方程；若未知函数是多元函数，则称为偏微分方程。
- **微分方程的阶**：微分方程中出现的未知函数的最高阶导数的阶数称为微分方程的阶。
- **解、通解、特解**：满足微分方程的函数称为微分方程的解。对于$n$阶微分方程，含有$n$个任意常数的解称为通解。通过给定的初始条件确定了通解中的任意常数后得到的解称为特解。

#### 一阶微分方程 📊

- **可分离变量方程**：形如$\frac{dy}{dx} = \frac{f(x)}{g(y)}$的方程，可通过分离变量法求解，即改写为$g(y) dy = f(x) dx$，然后两边积分得到通解。
- **齐次方程**：形如$\frac{dy}{dx} = \varphi\left( \frac{y}{x} \right)$的方程，可通过变量替换$u = \frac{y}{x}$转化为可分离变量方程求解。
- **一阶线性微分方程**：
  - **齐次线性方程**：形如$\frac{dy}{dx} + P(x) y = 0$的方程，其通解为
    $$
    y = C e^{- \int P(x) dx},
    $$
    其中$C$为任意常数。
  - **非齐次线性方程**：形如$\frac{dy}{dx} + P(x) y = Q(x)$的方程，其通解可通过常数变易法求得，即先求出对应的齐次方程的通解$y_h = C e^{- \int P(x) dx}$，然后将$C$视为$x$的函数$C(x)$，代入原方程求解$C(x)$，得到通解为
    $$
    y = e^{- \int P(x) dx} \left( C + \int Q(x) e^{\int P(x) dx} dx \right).
    $$
    或者，通解可表示为
    $$
    y = e^{- \int P(x) dx} \left( \int Q(x) e^{\int P(x) dx} dx + C \right).
    $$
- **伯努利方程**：形如$\frac{dy}{dx} + P(x) y = Q(x) y^n$的方程，可通过变量替换$v = y^{1 - n}$转化为线性微分方程求解。

#### 高阶微分方程 📊

- **可降阶的高阶微分方程**：
  - **方程$y^{(n)} = f(x)$**：通过$n$次积分可求得通解。
  - **方程$y'' = f(x, y')$**：令$p = y'$，则方程化为$p' = f(x, p)$，这是一个关于$p$和$x$的一阶方程，求出$p$后再积分得到$y$。
  - **方程$y'' = f(y, y')$**：令$p = y'$，则$p' = f(y, p)$，这是一个关于$p$和$y$的一阶方程，可分离变量求解，再积分得到$y$。
- **二阶线性微分方程**：
  - **齐次方程**：形如$y'' + P(x) y' + Q(x) y = 0$的方程。若已知一个特解$y_1$，可通过降阶法求得另一个线性无关的特解$y_2$，从而得到通解$y = C_1 y_1 + C_2 y_2$。
  - **常系数齐次线性方程**：形如$y'' + p y' + q y = 0$（$p$、$q$为常数）的方程，其特征方程为$r^2 + p r + q = 0$。根据特征根的不同情况，通解形式如下：
    - **两个不相等的实根$r_1$、$r_2$**：
      $$
      y = C_1 e^{r_1 x} + C_2 e^{r_2 x}.
      $$
    - **重根$r_1 = r_2 = r$**：
      $$
      y = (C_1 + C_2 x) e^{r x}.
      $$
    - **共轭复根$\alpha \pm i \beta$**：
      $$
      y = e^{\alpha x} (C_1 \cos \beta x + C_2 \sin \beta x).
      $$
  - **非齐次方程**：形如$y'' + p y' + q y = f(x)$的方程，其通解为对应的齐次方程的通解加上非齐次方程的一个特解$y_p$，即$y = y_h + y_p$。求特解$y_p$的方法有：
    - **待定系数法**：适用于$f(x)$为多项式、指数函数、正弦函数、余弦函数等简单函数的情况。根据$f(x)$的形式假设特解的形式，代入方程确定系数。
    - **常数变易法**：对于一般的非齐次方程，设特解为$y_p = u_1(x) y_1(x) + u_2(x) y_2(x)$，其中$y_1$、$y_2$是齐次方程的两个线性无关解，$u_1' y_1 + u_2' y_2 = 0$，$u_1' y_1' + u_2' y_2' = f(x)$，解此方程组可得$u_1'$、$u_2'$，积分后得到$u_1$、$u_2$，从而得到特解。