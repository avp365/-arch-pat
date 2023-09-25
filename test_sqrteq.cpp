#define BOOST_TEST_MODULE sqrteq_test_module

#include "sqrteq.h"

#include <boost/test/included/unit_test.hpp>

BOOST_AUTO_TEST_SUITE(sqrteq_test_suite)

// Написать тест, который проверяет, что для уравнения x^2+1 = 0 корней нет (возвращается пустой массив)
BOOST_AUTO_TEST_CASE(no_solution_of_an_equation)
{
    Sqrteq s;

    BOOST_CHECK(s.solve(1, 0, 1).empty() == 1);
}
// Написать тест, который проверяет, что для уравнения x^2-1 = 0 есть два корня кратности 1 (x1=1, x2=-1)
BOOST_AUTO_TEST_CASE(solution_1)
{
    Sqrteq s;
    std::vector<double> result;

    result = s.solve(1, 0, -1);

    BOOST_CHECK((result[0] == 1) && (result[1] == -1));
}

// Написать тест, который проверяет, что для уравнения x^2+2x+1 = 0 есть один корень кратности 2 (x1= x2 = -1).
BOOST_AUTO_TEST_CASE(solution_2)
{
    Sqrteq s;
    std::vector<double> result;

    result = s.solve(1, 2, 1);

    BOOST_CHECK((result[0] == -1) && (result[1] == -1));
}
// Написать тест, который проверяет, что коэффициент a не может быть равен 0. В этом случае solve выбрасывает исключение.
BOOST_AUTO_TEST_CASE(a_not_equal_zero)
{
    Sqrteq s;
    std::vector<double> result;

    BOOST_CHECK_THROW(s.solve(0.0000000456, 2, 1), std::invalid_argument);
}

// С учетом того, что дискриминант тоже нельзя сравнивать с 0 через знак равенства, подобрать
// такие коэффициенты квадратного уравнения для случая одного корня кратности два, чтобы
// дискриминант был отличный от нуля, но меньше заданного эпсилон.

BOOST_AUTO_TEST_CASE(d_less_e)
{
    Sqrteq s;
    std::vector<double> result;

    BOOST_CHECK(s.solve(0.0001, 0.0001, 0.0001).empty() == 0);
}

// Проверка на бесконечность a

BOOST_AUTO_TEST_CASE(test_infinity_a_b_c)
{
    Sqrteq s;
    std::vector<double> result;

    BOOST_CHECK(s.solve(INFINITY, 0.0001, 0.0001).empty() == 1);
    BOOST_CHECK(s.solve(-INFINITY, 0.0001, 0.0001).empty() == 1);

    BOOST_CHECK(s.solve(21, INFINITY, 0.0001).empty() == 1);
    BOOST_CHECK(s.solve(21, -INFINITY, 0.0001).empty() == 1);

    BOOST_CHECK(s.solve(33, 0.0001, INFINITY).empty() == 1);
    BOOST_CHECK(s.solve(33, 0.0001, -INFINITY).empty() == 1);
}

// Проверяет b и с чтобы так не были меньше e
BOOST_AUTO_TEST_CASE(b_no_more_max)
{
    Sqrteq s;
    std::vector<double> result;

    BOOST_CHECK_THROW(s.solve(34, 10000000230, 1), std::invalid_argument);
}

BOOST_AUTO_TEST_SUITE_END()