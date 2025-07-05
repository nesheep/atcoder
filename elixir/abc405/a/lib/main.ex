defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([r, x]), do: {String.to_integer(r), String.to_integer(x)}

  defp solve({r, 1}), do: r >= 1600 && r < 3000
  defp solve({r, 2}), do: r >= 1200 && r < 2400

  defp output(true), do: "Yes"
  defp output(false), do: "No"
end
