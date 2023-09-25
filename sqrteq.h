#pragma once
#include <vector>

class Sqrteq
{
private:
    double e = 1E-5;
    double max = 1000000000;

public:
    std::vector<double> solve(double a, double b, double c);
    double eq(double a, double b, double c, double x);
};