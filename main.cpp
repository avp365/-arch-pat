#include "sqrteq.h"
#include <iostream>
#include <math.h>

int main()
{
    Sqrteq s;
    double a, b, c;

    std::vector<double> res;

    a = INFINITY;
    b = 2;
    c = 1;
    std::cout << "test " << a << "\n";

    std::cout << "a " << a << "\n";
    std::cout << "b " << b << "\n";
    std::cout << "c " << c << "\n";
    std::cout
        << "start "
        << "\n";

    res = s.solve(a, b, c);

    std::cout << "solution " << res.size() << "\n";

    if (res.size() == 1)
    {
        std::cout << "x1 " << res[0] << "\n";
        std::cout << "res1 " << s.eq(a, b, c, res[0]) << "\n";
    }

    if (res.size() == 2)
    {
        std::cout << "x1 " << res[0] << "\n";
        std::cout << "x2 " << res[1] << "\n";

        std::cout << "res1 " << s.eq(a, b, c, res[0]) << "\n";
        std::cout << "res2 " << s.eq(a, b, c, res[1]) << "\n";
    }

    std::cout << "finish ";

    return 0;
}