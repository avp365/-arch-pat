#include "sqrteq.h"
#include <vector>
#include <math.h>
#include <iostream>

std::vector<double> Sqrteq::solve(double a, double b, double c)
{

    if (std::isinf(a) == 1 || std::isinf(b) == 1 || std::isinf(c) == 1)
    {

        return {};
    }

    std::cout << "a " << a << "\n";
    std::cout << "b " << b << "\n";
    std::cout << "c " << c << "\n";

    if (abs(a) <= Sqrteq::e)
    {
        throw std::invalid_argument("a b, c не должно быть меньше " + std::to_string(Sqrteq::e));
    }

    if (abs(b) > Sqrteq::max)
    {
        throw std::invalid_argument("b не больше  " + std::to_string(Sqrteq::max));
    }

    double x1, x2, D;

    D = b * b - 4 * a * c;

    std::cout << "D " << D << "\n";

    if (D < -Sqrteq::e)
    {
        return {};
    }

    if (abs(D) <= Sqrteq::e)
    {
        x1 = (-1 * b) / (2 * a);

        return {x1, x1};
    }
    if (D > Sqrteq::e)
    {
        x1 = (-1 * b + sqrt(D)) / (2 * a);

        x2 = (-1 * b - sqrt(D)) / (2 * a);

        return {x1, x2};
    }

    return {};
}
double Sqrteq::eq(double a, double b, double c, double x)
{
    return a * pow(x, 2) + b * x + c;
}