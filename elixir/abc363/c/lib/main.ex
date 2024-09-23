defmodule Main do
  def main do
    [n, k] = read_ints()
    s = IO.read(:line) |> String.trim()

    ans =
      s
      |> String.graphemes()
      |> Enum.sort()
      |> distinct_permutations()
      |> Stream.map(&Enum.join(&1))
      |> Stream.map(fn t -> 0..(n - k) |> Stream.map(&String.slice(t, &1, k)) end)
      |> Stream.map(fn l -> l |> Stream.map(&(&1 != String.reverse(&1))) end)
      |> Stream.map(&Enum.all?(&1))
      |> Enum.count(& &1)

    IO.puts(ans)
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  # Algorithm: https://w.wiki/Qai
  def distinct_permutations(s) do
    step1 = fn t ->
      t
      |> Enum.chunk_every(2, 1, :discard)
      |> Enum.find(fn [{b, _}, {a, _}] -> a < b end)
      |> then(fn c ->
        case c do
          [{_, _}, {_, _} = a] -> a
          _ -> throw(:exit)
        end
      end)
    end

    step2 = fn t, a ->
      t
      |> Enum.find(fn {b, _} -> a < b end)
      |> then(fn c ->
        case c do
          {_, _} = b -> b
          _ -> throw(:exit)
        end
      end)
    end

    step3 = fn s, {a, k}, {b, l} ->
      s
      |> List.replace_at(k, b)
      |> List.replace_at(l, a)
    end

    step4 = fn u, k ->
      (u |> Enum.slice(0..k)) ++ (u |> Enum.slice((k + 1)..-1//1) |> Enum.reverse())
    end

    Stream.iterate({:cont, s}, fn {_, s} ->
      try do
        t = s |> Enum.with_index() |> Enum.reverse()
        {a, k} = step1.(t)
        {b, l} = step2.(t, a)
        u = s |> step3.({a, k}, {b, l}) |> step4.(k)
        {:cont, u}
      catch
        :exit -> {:halt, []}
      end
    end)
    |> Stream.take_while(fn {c, _} -> c == :cont end)
    |> Stream.map(fn {_, s} -> s end)
  end
end
