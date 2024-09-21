defmodule Main do
  def main do
    [xa, ya] = read_ints()
    [xb, yb] = read_ints()
    [xc, yc] = read_ints()

    ab = square(xb - xa) + square(yb - ya)
    bc = square(xc - xb) + square(yc - yb)
    ca = square(xa - xc) + square(ya - yc)

    m = Enum.max([ab, bc, ca])
    ans = m == Enum.sum(Enum.reject([ab, bc, ca], &(&1 == m)))

    IO.puts(if ans, do: "Yes", else: "No")
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  def square(x), do: x * x
end
